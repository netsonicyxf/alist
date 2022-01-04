package bootstrap

import (
	"github.com/Xhofe/alist/conf"
	"github.com/Xhofe/alist/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
)

func InitSettings() {
	log.Infof("init settings...")

	err := model.SaveSetting(model.Version)
	if err != nil {
		log.Fatalf("failed write setting: %s", err.Error())
	}

	settings := []model.SettingItem{
		{
			Key:         "title",
			Value:       "Alist",
			Description: "title",
			Type:        "string",
			Access:      model.PUBLIC,
			Group:       model.FRONT,
		},
		{
			Key:         "password",
			Value:       "alist",
			Description: "password",
			Type:        "string",
			Access:      model.PRIVATE,
			Group:       model.BACK,
		},
		{
			Key:         "logo",
			Value:       "https://store.heytapimage.com/cdo-portal/feedback/202112/05/1542f45f86b8609495b69c5380753135.png",
			Description: "logo",
			Type:        "string",
			Access:      model.PUBLIC,
			Group:       model.FRONT,
		},
		{
			Key:         "favicon",
			Value:       "https://store.heytapimage.com/cdo-portal/feedback/202112/05/1542f45f86b8609495b69c5380753135.png",
			Description: "favicon",
			Type:        "string",
			Access:      model.PUBLIC,
			Group:       model.FRONT,
		},
		{
			Key:         "icon color",
			Value:       "#1890ff",
			Description: "icon's color",
			Type:        "string",
			Access:      model.PUBLIC,
			Group:       model.FRONT,
		},
		{
			Key:         "text types",
			Value:       strings.Join(conf.TextTypes, ","),
			Type:        "string",
			Description: "text type extensions",
			Group:       model.FRONT,
		},
		{
			Key:         "hide readme file",
			Value:       "true",
			Type:        "bool",
			Description: "hide readme file? ",
			Group:       model.FRONT,
		},
		{
			Key:         "music cover",
			Value:       "https://store.heytapimage.com/cdo-portal/feedback/202110/30/d43c41c5d257c9bc36366e310374fb19.png",
			Description: "music cover image",
			Type:        "string",
			Access:      model.PUBLIC,
			Group:       model.FRONT,
		},
		{
			Key:         "site beian",
			Description: "chinese beian info",
			Type:        "string",
			Access:      model.PUBLIC,
			Group:       model.FRONT,
		},
		{
			Key:         "home readme url",
			Description: "when have multiple, the readme file to show",
			Type:        "string",
			Access:      model.PUBLIC,
			Group:       model.FRONT,
		},
		{
			Key:    "autoplay video",
			Value:  "false",
			Type:   "bool",
			Access: model.PUBLIC,
			Group:  model.FRONT,
		},
		{
			Key:    "autoplay audio",
			Value:  "false",
			Type:   "bool",
			Access: model.PUBLIC,
			Group:  model.FRONT,
		},
		{
			Key:         "check parent folder",
			Value:       "false",
			Type:        "bool",
			Description: "check parent folder password",
			Access:      model.PRIVATE,
			Group:       model.BACK,
		},
		{
			Key:         "customize head",
			Value:       "",
			Type:        "text",
			Description: "Customize head, placed at the beginning of the head",
			Access:      model.PRIVATE,
			Group:       model.FRONT,
		},
		{
			Key:         "customize body",
			Value:       "",
			Type:        "text",
			Description: "Customize script, placed at the end of the body",
			Access:      model.PRIVATE,
			Group:       model.FRONT,
		},
		{
			Key:         "home emoji",
			Value:       "🏠",
			Type:        "string",
			Description: "emoji in front of home in nav",
			Access:      model.PUBLIC,
			Group:       model.FRONT,
		},
		{
			Key:         "animation",
			Value:       "true",
			Type:        "bool",
			Description: "when there are a lot of files, the animation will freeze when opening",
			Access:      model.PUBLIC,
			Group:       model.FRONT,
		},
		{
			Key:         "check down link",
			Value:       "false",
			Type:        "bool",
			Description: "check down link password, your link will be 'https://alist.com/d/filename?pw=xxx'",
			Access:      model.PUBLIC,
			Group:       model.BACK,
		},
		{
			Key:         "WebDAV username",
			Value:       "alist_admin",
			Description: "WebDAV username",
			Type:        "string",
			Access:      model.PRIVATE,
			Group:       model.BACK,
		},
		{
			Key:         "WebDAV password",
			Value:       "alist_admin",
			Description: "WebDAV password",
			Type:        "string",
			Access:      model.PRIVATE,
			Group:       model.BACK,
		},
		{
			Key:         "artplayer whitelist",
			Value:       "*",
			Description: "refer to https://artplayer.org/document/options#whitelist",
			Type:        "string",
			Access:      model.PUBLIC,
			Group:       model.FRONT,
		},
		{
			Key:         "artplayer autoSize",
			Value:       "true",
			Description: "refer to https://artplayer.org/document/options#autosize",
			Type:        "bool",
			Access:      model.PUBLIC,
			Group:       model.FRONT,
		},
		{
			Key:         "Visitor WebDAV username",
			Value:       "alist_visitor",
			Description: "Visitor WebDAV username",
			Type:        "string",
			Access:      model.PRIVATE,
			Group:       model.BACK,
		},
		{
			Key:         "Visitor WebDAV password",
			Value:       "alist_visitor",
			Description: "Visitor WebDAV password",
			Type:        "string",
			Access:      model.PRIVATE,
			Group:       model.BACK,
		},
		{
			Key:    "load type",
			Value:  "all",
			Type:   "select",
			Values: "all,load more,auto load more,pagination",
			Access: model.PUBLIC,
			Group:  model.FRONT,
		},
		{
			Key:    "default page size",
			Value:  "30",
			Type:   "number",
			Access: model.PUBLIC,
			Group:  model.FRONT,
		},
	}
	for i, _ := range settings {
		v := settings[i]
		v.Version = conf.GitTag
		o, err := model.GetSettingByKey(v.Key)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				err = model.SaveSetting(v)
				if err != nil {
					log.Fatalf("failed write setting: %s", err.Error())
				}
			} else {
				log.Fatal("can't get setting: %s", err.Error())
			}
		} else {
			//o.Version = conf.GitTag
			//err = model.SaveSetting(*o)
			v.Value = o.Value
			err = model.SaveSetting(v)
			if err != nil {
				log.Fatalf("failed write setting: %s", err.Error())
			}
		}
	}
	model.LoadSettings()
}
