package main

import _ "net/http/pprof"
import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/adrg/xdg"
	"github.com/alessio/shellescape"
	"github.com/bmmcginty/barnard/config"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/bmmcginty/barnard/gumble/gumble"
	_ "github.com/bmmcginty/barnard/gumble/opus"
	"github.com/bmmcginty/barnard/uiterm"
	"github.com/bmmcginty/go-openal/openal"
)

func show_devs(name string, args []string) {
	if args == nil {
		fmt.Printf("no items for %s\n", name)
	}
	fmt.Printf("%s\n", name)
	for i := 0; i < len(args); i++ {
		fmt.Printf("%s\n", args[i])
	}
}

func do_list_devices() {
	odevs := openal.GetStrings(openal.AllDevicesSpecifier)
	if odevs != nil && len(odevs) > 0 {
		show_devs("All outputs:", odevs)
	} else {
		odevs = openal.GetStrings(openal.DeviceSpecifier)
		show_devs("All outputs:", odevs)
	}
	idevs := openal.GetStrings(openal.CaptureDeviceSpecifier)
	show_devs("Inputs:", idevs)
}

func setup_notify_runner(notify_command string) chan []string {
	t := make(chan []string)
	var do_nothing = false
	var err error
	if err != nil {
	}
	if notify_command == "" {
		do_nothing = true
	}
	go func(events chan []string, cmd_template string, dummy bool) {
		for {
			event := <-events
			if !dummy {
				t := string(cmd_template)
				t = strings.ReplaceAll(t, "%event", shellescape.Quote(event[0]))
				t = strings.ReplaceAll(t, "%who", shellescape.Quote(event[1]))
				t = strings.ReplaceAll(t, "%what", shellescape.Quote(event[2]))
				cmd := "/bin/sh"
				args := []string{"-c", t}
				x := exec.Command(cmd, args...)
				x.Run()
			} //if we actually have a command to run
		} //for
	}(t, notify_command, do_nothing)
	return t
}

func setup_fifo(fn string) (chan string, error) {
	t := make(chan string)
	if fn == "" {
		return t, nil
	}
	os.Remove(fn)
	err := syscall.Mkfifo(fn, 0600)
	if err != nil {
		return t, err
	}
	file, err := os.OpenFile(fn, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		return t, err
	}
	go func(fh io.Reader, out chan string) {
		reader := bufio.NewReader(fh)
		for {
			line, err := reader.ReadBytes('\n')
			if err == nil {
				out <- strings.TrimSpace(string(line))
			}
		}
	}(file, t)
	return t, nil
}

func main() {
	configFilePath, err := xdg.ConfigFile("barnard/config.yaml")
	if err != nil {
		configFilePath = "~/.barnard.yaml"
		log.Fatal(err)
	}

	// Command line flags
	server := flag.String("server", "localhost:64738", "the server to connect to")
	username := flag.String("username", "", "the username of the client")
	password := flag.String("password", "", "the password of the server")
	insecure := flag.Bool("insecure", false, "skip server certificate verification")
	certificate := flag.String("certificate", "", "PEM encoded certificate and private key")
	cfgfn := flag.String("config", configFilePath, "Path to YAML formatted configuration file")
	list_devices := flag.Bool("list_devices", false, "do not connect; instead, list available audio devices and exit")
	fifo := flag.String("fifo", "", "path of a FIFO from which to read commands")
	serverSet := false
	usernameSet := false
	buffers := flag.Int("buffers", 16, "number of audio buffers to use")
	profile := flag.Bool("profile", false, "add http server to serve profiles")

	flag.Parse()

	if *profile == true {
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	}

	userConfig := config.NewConfig(cfgfn)

	flag.CommandLine.Visit(func(theFlag *flag.Flag) {
		switch theFlag.Name {
		case "server":
			serverSet = true
		case "username":
			usernameSet = true
		}
	})

	if !serverSet {
		server = userConfig.GetDefaultServer()
	}
	if !usernameSet {
		username = userConfig.GetUsername()
	}

	if os.Getenv("ALSOFT_LOGLEVEL") == "" {
		os.Setenv("ALSOFT_LOGLEVEL", "0")
	}

	if *list_devices {
		do_list_devices()
		os.Exit(0)
	}

	if !strings.Contains(*server, ":") {
		*server = (*server + ":64738")
	}

	// Initialize
	b := Barnard{
		Config:     gumble.NewConfig(),
		UserConfig: userConfig,
		Address:    *server,
	}
	b.Config.Buffers = *buffers

	b.Hotkeys = b.UserConfig.GetHotkeys()
	b.UserConfig.SaveConfig()
	b.Config.Username = *username
	b.Config.Password = *password

	if *insecure {
		b.TLSConfig.InsecureSkipVerify = true
	}
	if *certificate != "" {
		cert, err := tls.LoadX509KeyPair(*certificate, *certificate)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		b.TLSConfig.Certificates = append(b.TLSConfig.Certificates, cert)
	}

	reader, err := setup_fifo(*fifo)
	if err != nil {
		b.exitMessage = err.Error()
		b.exitStatus = 1
		handle_error(b)
	}
	b.notifyChannel = setup_notify_runner(*b.UserConfig.GetNotifyCommand())
	b.Ui = uiterm.New(&b)
	b.Ui.Run(reader)
	handle_error(b)
}

func handle_raw_error(e error) {
	fmt.Fprintf(os.Stderr, "%s\n", e.Error())
	os.Exit(1)
}

func handle_error(b Barnard) {
	if b.exitMessage != "" {
		fmt.Fprintf(os.Stderr, "%s\n", b.exitMessage)
	}
	os.Exit(b.exitStatus)
}
