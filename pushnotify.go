package mob_push_sdk

type PushNotify struct {
	TaskCron       int            `json:"taskCron"`
	TaskTime       int64          `json:"taskTime"`
	Plats          []int          `json:"plats"`
	IosProduction  int            `json:"iosProduction"`
	OfflineSeconds int            `json:"offlineSeconds"`
	Content        string         `json:"content"`
	Title          string         `json:"title"`
	Type           int            `json:"type"`
	Url            string         `json:"url"`
	CustomNotify   *CustomNotify  `json:"customNotify"`
	AndroidNotify  *AndroidNotify `json:"androidNotify"`
	IosNotify      *IosNotify     `json:"iosNotify"`
	ExtrasMapList  []PushMap      `json:"extrasMapList"`
	Speed          int            `json:"speed"`
}

type CustomNotify struct {
	CustomType  string `json:"customType"`
	CustomTitle string `json:"customTitle"`
}

type AndroidNotify struct {
	AppName   string   `json:"appName,omitempty"`
	Title     string   `json:"title,omitempty"`
	Warn      string   `json:"warn,omitempty"`
	Style     int      `json:"style,omitempty"`
	Content   []string `json:"content,omitempty"`
	Sound     string   `json:"sound,omitempty"`
	Badge     int      `json:"androidBadge,omitempty"`
	BadgeType int      `json:"androidBadgeType,omitempty"`
}

type IosNotify struct {
	Badge            int    `json:"badge,omitempty"`
	BadgeType        int    `json:"badgeType,omitempty"`
	Category         string `json:"category,omitempty"`
	Sound            string `json:"sound,omitempty"`
	SubTitle         string `json:"subtitle,omitempty"`
	SlientPush       int    `json:"slientPush,omitempty"`
	ContentAvailable int    `json:"contentAvailable,omitempty"`
	MutableContent   int    `json:"mutableContent,omitempty"`
	AttachmentType   int    `json:"attachmentType,omitempty"`
	Attachment       string `json:"attachment,omitempty"`
}

type PushMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (push *Push) setTitle(title string) *Push {
	push.PushNotify.Title = title
	return push
}

func (push *Push) setContent(content string) *Push {
	push.PushNotify.Content = content
	return push
}

func (push *Push) setPlats(plats []int) *Push {
	push.PushNotify.Plats = plats
	return push
}

func (push *Push) setCustomNotify(customNotify CustomNotify) *Push {
	push.PushNotify.CustomNotify = &customNotify
	return push
}

func (push *Push) setAndroidNotify(androidNotify AndroidNotify) *Push {
	push.PushNotify.AndroidNotify = &androidNotify
	return push
}

func (push *Push) setIosNotify(iosNotify IosNotify) *Push {
	push.PushNotify.IosNotify = &iosNotify
	return push
}

func (push *Push) SetIOSProduction(iosProduction int) *Push {
	push.PushNotify.IosProduction = iosProduction
	return push
}
