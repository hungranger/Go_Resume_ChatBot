package config

import (
	"reflect"
	"testing"
)

func TestIniConfig_LoadConfig(t *testing.T) {
	c := new(IniConfig)

	got := c.LoadConfig()
	want := &IniConfig{
		Port:         "8080",
		Facebook_api: "https://graph.facebook.com/v2.6/me/messages?access_token=%s",
		Image:        "http://37.media.tumblr.com/e705e901302b5925ffb2bcf3cacb5bcd/tumblr_n6vxziSQD11slv6upo3_500.gif",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("IniConfig.LoadConfig() = %v, want %v", got, want)
	}
}
