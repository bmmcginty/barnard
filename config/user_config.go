package config

import (
	"fmt"
	"github.com/bmmcginty/barnard/uiterm"
	"gopkg.in/yaml.v2"
	"github.com/bmmcginty/barnard/gumble/gumble"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"strings"
)

type Config struct {
	config *exportableConfig
	fn     string
}

type exportableConfig struct {
	Hotkeys       *Hotkeys
	MicVolume     *float32
	InputDevice   *string
	OutputDevice  *string
	Servers       []*server
	DefaultServer *string
	Username      *string
	NotifyCommand *string
}

type server struct {
	Host  string
	Port  int
	Users []*eUser
}

type eUser struct {
	Username string
	Boost    uint16
	Volume   float32
}

func (c *Config) SaveConfig() {
	var data []byte
	data, err := yaml.Marshal(c.config)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(c.fn+".tmp", data, 0600)
	if err != nil {
		panic(err)
	}
	err = os.Rename(c.fn+".tmp", c.fn)
	if err != nil {
		panic(err)
	}
}

func key(k uiterm.Key) *uiterm.Key {
	return &k
}

func (c *Config) LoadConfig() {
	var jc exportableConfig
	jc = exportableConfig{}
	jc.Hotkeys = &Hotkeys{
		Talk:             key(uiterm.KeyF1),
		VolumeDown:       key(uiterm.KeyF5),
		VolumeUp:         key(uiterm.KeyF6),
		Exit:             key(uiterm.KeyF10),
		ToggleTimestamps: key(uiterm.KeyF3),
		SwitchViews:      key(uiterm.KeyTab),
		ScrollUp:         key(uiterm.KeyPgup),
		ScrollDown:       key(uiterm.KeyPgdn),
	}
	if fileExists(c.fn) {
		var data []byte
		data = readFile(c.fn)
		if data != nil {
			err := yaml.UnmarshalStrict(data, &jc)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing \"%s\".\n%s\n", c.fn, err.Error())
				os.Exit(1)
			} //panic
		} //if data
	} //if exists
	c.config = &jc
	if c.config.MicVolume == nil {
		micvol := float32(1.0)
		jc.MicVolume = &micvol
	}
	if c.config.InputDevice == nil {
		idev := string("")
		jc.InputDevice = &idev
	}
	if c.config.OutputDevice == nil {
		odev := string("")
		jc.OutputDevice = &odev
	}
	if c.config.DefaultServer == nil {
		defaultServer := string("localhost:64738")
		jc.DefaultServer = &defaultServer
	}
	if c.config.Username == nil {
		username := string("")
		jc.Username = &username
	}
	if c.config.NotifyCommand == nil {
		ncmd := string("")
		jc.NotifyCommand = &ncmd
	}
}

func (c *Config) findServer(address string) *server {
	if c.config.Servers == nil {
		c.config.Servers = make([]*server, 0)
	}
	host, port := makeHostPort(address)
	var t *server
	for _, s := range c.config.Servers {
		if s.Port == port && s.Host == host {
			t = s
			break
		}
	}
	if t == nil {
		t = &server{
			Host: host,
			Port: port,
		}
		c.config.Servers = append(c.config.Servers, t)
	}
	return t
}

func (c *Config) findUser(address string, username string) *eUser {
	var s *server
	s = c.findServer(address)
	if s.Users == nil {
		s.Users = make([]*eUser, 0)
	}
	var t *eUser
	for _, u := range s.Users {
		if u.Username == username {
			t = u
			break
		}
	}
	if t == nil {
		t = &eUser{
			Username: username,
			Boost:    uint16(1),
			Volume:   1.0,
		}
		s.Users = append(s.Users, t)
	}
	return t
}

func (c *Config) SetMicVolume(v float32) {
	t := float32(v)
	c.config.MicVolume = &t
}

func (c *Config) GetHotkeys() *Hotkeys {
	return c.config.Hotkeys
}

func (c *Config) GetNotifyCommand() *string {
	return c.config.NotifyCommand
}

func (c *Config) GetInputDevice() *string {
	return c.config.InputDevice
}

func (c *Config) GetOutputDevice() *string {
	return c.config.OutputDevice
}

func (c *Config) GetDefaultServer() *string {
	return c.config.DefaultServer
}

func (c *Config) GetUsername() *string {
	return c.config.Username
}

func (c *Config) UpdateUser(u *gumble.User) {
	var j *eUser
	var uc *gumble.Client
	uc = u.GetClient()
	if uc != nil {
		j = c.findUser(uc.Config.Address, u.Name)
		u.Boost = j.Boost
		u.Volume = j.Volume
		if u.Boost < 1 {
			u.Boost = 1
		}
	}
}

func (c *Config) UpdateConfig(u *gumble.User) {
	var j *eUser
	j = c.findUser(u.GetClient().Config.Address, u.Name)
	j.Boost = u.Boost
	j.Volume = u.Volume
}

func NewConfig(fn *string) *Config {
	var c *Config
	c = &Config{}
	c.fn = resolvePath(*fn)
	c.LoadConfig()
	return c
}

func readFile(path string) []byte {
	if !fileExists(path) {
		return nil
	}
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return dat
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func resolvePath(path string) string {
	if strings.HasPrefix(path, "~/") {
		usr, err := user.Current()
		if err != nil {
			panic(err)
		}
    path = strings.Replace(path, "~", usr.HomeDir, 1)
  }

	return path
}

func makeHostPort(addr string) (string, int) {
	parts := strings.Split(addr, ":")
	host := parts[0]
	port, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return host, port
}

func Log(s string) {
	log(s)
}

func log(s string) {
	s += "\n"
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	if _, err := f.Write([]byte(s)); err != nil {
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}
}
