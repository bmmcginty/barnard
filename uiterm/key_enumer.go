// Code generated by "enumer -type=Key -trimprefix=Key -yaml -json -transform=snake"; DO NOT EDIT.

//
package uiterm

import (
	"encoding/json"
	"fmt"
)

const (
	_KeyName_0 = "ctrl_tildectrl_actrl_bctrl_cctrl_dctrl_ectrl_fctrl_gbackspacetabctrl_jctrl_kctrl_lenterctrl_nctrl_octrl_pctrl_qctrl_rctrl_sctrl_tctrl_uctrl_vctrl_wctrl_xctrl_yctrl_zescctrl4ctrl5ctrl6ctrl7space"
	_KeyName_1 = "backspace2"
	_KeyName_2 = "mouse_rightmouse_middlemouse_leftarrow_rightarrow_leftarrow_downarrow_uppgdnpgupendhomedeleteinsertf12f11f10f9f8f7f6f5f4f3f2f1alt_ctrl_tildealt_ctrl_aalt_ctrl_balt_ctrl_calt_ctrl_dalt_ctrl_ealt_ctrl_falt_ctrl_galt_backspacealt_tabalt_ctrl_jalt_ctrl_kalt_ctrl_lalt_enteralt_ctrl_nalt_ctrl_oalt_ctrl_palt_ctrl_qalt_ctrl_ralt_ctrl_salt_ctrl_talt_ctrl_ualt_ctrl_valt_ctrl_walt_ctrl_xalt_ctrl_yalt_ctrl_zalt_escalt_ctrl4alt_ctrl5alt_ctrl6alt_ctrl7alt_space"
	_KeyName_3 = "alt_aalt_balt_calt_dalt_ealt_falt_galt_halt_ialt_jalt_kalt_lalt_malt_nalt_oalt_palt_qalt_ralt_salt_talt_ualt_valt_walt_xalt_yalt_z"
	_KeyName_4 = "alt_backspace2"
	_KeyName_5 = "alt_arrow_rightalt_arrow_leftalt_arrow_downalt_arrow_upalt_pgdnalt_pgupalt_endalt_homealt_deletealt_insertalt_f12alt_f11alt_f10alt_f9alt_f8alt_f7alt_f6alt_f5alt_f4alt_f3alt_f2alt_f1"
)

var (
	_KeyIndex_0 = [...]uint8{0, 10, 16, 22, 28, 34, 40, 46, 52, 61, 64, 70, 76, 82, 87, 93, 99, 105, 111, 117, 123, 129, 135, 141, 147, 153, 159, 165, 168, 173, 178, 183, 188, 193}
	_KeyIndex_1 = [...]uint8{0, 10}
	_KeyIndex_2 = [...]uint16{0, 11, 23, 33, 44, 54, 64, 72, 76, 80, 83, 87, 93, 99, 102, 105, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 140, 150, 160, 170, 180, 190, 200, 210, 223, 230, 240, 250, 260, 269, 279, 289, 299, 309, 319, 329, 339, 349, 359, 369, 379, 389, 399, 406, 415, 424, 433, 442, 451}
	_KeyIndex_3 = [...]uint8{0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100, 105, 110, 115, 120, 125, 130}
	_KeyIndex_4 = [...]uint8{0, 14}
	_KeyIndex_5 = [...]uint8{0, 15, 29, 43, 55, 63, 71, 78, 86, 96, 106, 113, 120, 127, 133, 139, 145, 151, 157, 163, 169, 175, 181}
)

func (i Key) String() string {
	switch {
	case 0 <= i && i <= 32:
		return _KeyName_0[_KeyIndex_0[i]:_KeyIndex_0[i+1]]
	case i == 127:
		return _KeyName_1
	case 65511 <= i && i <= 65568:
		i -= 65511
		return _KeyName_2[_KeyIndex_2[i]:_KeyIndex_2[i+1]]
	case 65633 <= i && i <= 65658:
		i -= 65633
		return _KeyName_3[_KeyIndex_3[i]:_KeyIndex_3[i+1]]
	case i == 65663:
		return _KeyName_4
	case 131050 <= i && i <= 131071:
		i -= 131050
		return _KeyName_5[_KeyIndex_5[i]:_KeyIndex_5[i+1]]
	default:
		return fmt.Sprintf("Key(%d)", i)
	}
}

var _KeyValues = []Key{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 127, 65511, 65512, 65513, 65514, 65515, 65516, 65517, 65518, 65519, 65520, 65521, 65522, 65523, 65524, 65525, 65526, 65527, 65528, 65529, 65530, 65531, 65532, 65533, 65534, 65535, 65536, 65537, 65538, 65539, 65540, 65541, 65542, 65543, 65544, 65545, 65546, 65547, 65548, 65549, 65550, 65551, 65552, 65553, 65554, 65555, 65556, 65557, 65558, 65559, 65560, 65561, 65562, 65563, 65564, 65565, 65566, 65567, 65568, 65633, 65634, 65635, 65636, 65637, 65638, 65639, 65640, 65641, 65642, 65643, 65644, 65645, 65646, 65647, 65648, 65649, 65650, 65651, 65652, 65653, 65654, 65655, 65656, 65657, 65658, 65663, 131050, 131051, 131052, 131053, 131054, 131055, 131056, 131057, 131058, 131059, 131060, 131061, 131062, 131063, 131064, 131065, 131066, 131067, 131068, 131069, 131070, 131071}

var _KeyNameToValueMap = map[string]Key{
	_KeyName_0[0:10]:    0,
	_KeyName_0[10:16]:   1,
	_KeyName_0[16:22]:   2,
	_KeyName_0[22:28]:   3,
	_KeyName_0[28:34]:   4,
	_KeyName_0[34:40]:   5,
	_KeyName_0[40:46]:   6,
	_KeyName_0[46:52]:   7,
	_KeyName_0[52:61]:   8,
	_KeyName_0[61:64]:   9,
	_KeyName_0[64:70]:   10,
	_KeyName_0[70:76]:   11,
	_KeyName_0[76:82]:   12,
	_KeyName_0[82:87]:   13,
	_KeyName_0[87:93]:   14,
	_KeyName_0[93:99]:   15,
	_KeyName_0[99:105]:  16,
	_KeyName_0[105:111]: 17,
	_KeyName_0[111:117]: 18,
	_KeyName_0[117:123]: 19,
	_KeyName_0[123:129]: 20,
	_KeyName_0[129:135]: 21,
	_KeyName_0[135:141]: 22,
	_KeyName_0[141:147]: 23,
	_KeyName_0[147:153]: 24,
	_KeyName_0[153:159]: 25,
	_KeyName_0[159:165]: 26,
	_KeyName_0[165:168]: 27,
	_KeyName_0[168:173]: 28,
	_KeyName_0[173:178]: 29,
	_KeyName_0[178:183]: 30,
	_KeyName_0[183:188]: 31,
	_KeyName_0[188:193]: 32,
	_KeyName_1[0:10]:    127,
	_KeyName_2[0:11]:    65511,
	_KeyName_2[11:23]:   65512,
	_KeyName_2[23:33]:   65513,
	_KeyName_2[33:44]:   65514,
	_KeyName_2[44:54]:   65515,
	_KeyName_2[54:64]:   65516,
	_KeyName_2[64:72]:   65517,
	_KeyName_2[72:76]:   65518,
	_KeyName_2[76:80]:   65519,
	_KeyName_2[80:83]:   65520,
	_KeyName_2[83:87]:   65521,
	_KeyName_2[87:93]:   65522,
	_KeyName_2[93:99]:   65523,
	_KeyName_2[99:102]:  65524,
	_KeyName_2[102:105]: 65525,
	_KeyName_2[105:108]: 65526,
	_KeyName_2[108:110]: 65527,
	_KeyName_2[110:112]: 65528,
	_KeyName_2[112:114]: 65529,
	_KeyName_2[114:116]: 65530,
	_KeyName_2[116:118]: 65531,
	_KeyName_2[118:120]: 65532,
	_KeyName_2[120:122]: 65533,
	_KeyName_2[122:124]: 65534,
	_KeyName_2[124:126]: 65535,
	_KeyName_2[126:140]: 65536,
	_KeyName_2[140:150]: 65537,
	_KeyName_2[150:160]: 65538,
	_KeyName_2[160:170]: 65539,
	_KeyName_2[170:180]: 65540,
	_KeyName_2[180:190]: 65541,
	_KeyName_2[190:200]: 65542,
	_KeyName_2[200:210]: 65543,
	_KeyName_2[210:223]: 65544,
	_KeyName_2[223:230]: 65545,
	_KeyName_2[230:240]: 65546,
	_KeyName_2[240:250]: 65547,
	_KeyName_2[250:260]: 65548,
	_KeyName_2[260:269]: 65549,
	_KeyName_2[269:279]: 65550,
	_KeyName_2[279:289]: 65551,
	_KeyName_2[289:299]: 65552,
	_KeyName_2[299:309]: 65553,
	_KeyName_2[309:319]: 65554,
	_KeyName_2[319:329]: 65555,
	_KeyName_2[329:339]: 65556,
	_KeyName_2[339:349]: 65557,
	_KeyName_2[349:359]: 65558,
	_KeyName_2[359:369]: 65559,
	_KeyName_2[369:379]: 65560,
	_KeyName_2[379:389]: 65561,
	_KeyName_2[389:399]: 65562,
	_KeyName_2[399:406]: 65563,
	_KeyName_2[406:415]: 65564,
	_KeyName_2[415:424]: 65565,
	_KeyName_2[424:433]: 65566,
	_KeyName_2[433:442]: 65567,
	_KeyName_2[442:451]: 65568,
	_KeyName_3[0:5]:     65633,
	_KeyName_3[5:10]:    65634,
	_KeyName_3[10:15]:   65635,
	_KeyName_3[15:20]:   65636,
	_KeyName_3[20:25]:   65637,
	_KeyName_3[25:30]:   65638,
	_KeyName_3[30:35]:   65639,
	_KeyName_3[35:40]:   65640,
	_KeyName_3[40:45]:   65641,
	_KeyName_3[45:50]:   65642,
	_KeyName_3[50:55]:   65643,
	_KeyName_3[55:60]:   65644,
	_KeyName_3[60:65]:   65645,
	_KeyName_3[65:70]:   65646,
	_KeyName_3[70:75]:   65647,
	_KeyName_3[75:80]:   65648,
	_KeyName_3[80:85]:   65649,
	_KeyName_3[85:90]:   65650,
	_KeyName_3[90:95]:   65651,
	_KeyName_3[95:100]:  65652,
	_KeyName_3[100:105]: 65653,
	_KeyName_3[105:110]: 65654,
	_KeyName_3[110:115]: 65655,
	_KeyName_3[115:120]: 65656,
	_KeyName_3[120:125]: 65657,
	_KeyName_3[125:130]: 65658,
	_KeyName_4[0:14]:    65663,
	_KeyName_5[0:15]:    131050,
	_KeyName_5[15:29]:   131051,
	_KeyName_5[29:43]:   131052,
	_KeyName_5[43:55]:   131053,
	_KeyName_5[55:63]:   131054,
	_KeyName_5[63:71]:   131055,
	_KeyName_5[71:78]:   131056,
	_KeyName_5[78:86]:   131057,
	_KeyName_5[86:96]:   131058,
	_KeyName_5[96:106]:  131059,
	_KeyName_5[106:113]: 131060,
	_KeyName_5[113:120]: 131061,
	_KeyName_5[120:127]: 131062,
	_KeyName_5[127:133]: 131063,
	_KeyName_5[133:139]: 131064,
	_KeyName_5[139:145]: 131065,
	_KeyName_5[145:151]: 131066,
	_KeyName_5[151:157]: 131067,
	_KeyName_5[157:163]: 131068,
	_KeyName_5[163:169]: 131069,
	_KeyName_5[169:175]: 131070,
	_KeyName_5[175:181]: 131071,
}

// KeyString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func KeyString(s string) (Key, error) {
	if val, ok := _KeyNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Key values", s)
}

// KeyValues returns all values of the enum
func KeyValues() []Key {
	return _KeyValues
}

// IsAKey returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Key) IsAKey() bool {
	for _, v := range _KeyValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Key
func (i Key) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Key
func (i *Key) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Key should be a string, got %s", data)
	}

	var err error
	*i, err = KeyString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for Key
func (i Key) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Key
func (i *Key) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = KeyString(s)
	return err
}