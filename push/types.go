package push

type PushEntity struct {
	Platform        string        `json:"platform" comment:"oneOf android,ios,quickapp,hmos,all"`
	Audience        any           `json:"audience"`
	Notification    *Notification `json:"notification,omitempty"`
	Message         *Message      `json:"message,omitempty"`
	SmsMessage      *SmsMessage   `json:"sms_message,omitempty"`
	Options         *Options      `json:"options,omitempty"`
	Callback        *Callback     `json:"callback,omitempty"`
	Notification3rd any           `json:"notification_3rd,omitempty"`
	Cid             string        `json:"cid,omitempty"`
}

type Notification struct {
	Alert   string  `json:"alert,omitempty"`
	Android Android `json:"android"`
	Ios     Ios     `json:"ios"`
	Voip    any     `json:"voip,omitempty" comment:"ios透传"`
}

type Message struct {
	MsgContent  string `json:"msg_content"`
	ContentType string `json:"content_type,omitempty"`
	Title       string `json:"title,omitempty"`
	Extras      any    `json:"extras,omitempty"`
}

type Audience struct {
	Tag            []string `json:"tag,omitempty"`
	TagAnd         []string `json:"tag_and,omitempty"`
	TagNot         []string `json:"tag_not,omitempty"`
	Alias          []string `json:"alias,omitempty"`
	RegistrationId []string `json:"registration_id,omitempty"`
	Segment        []string `json:"segment,omitempty"`
	Abtest         []string `json:"abtest,omitempty"`
	LiveActivityID string   `json:"live_activity_id,omitempty"`
}

type Android struct {
	Alert             string         `json:"alert"`
	Title             string         `json:"title,omitempty"`
	BuilderId         int            `json:"builder_id,omitempty"`
	ChannelID         string         `json:"channel_id,omitempty"`
	Category          string         `json:"category,omitempty"`
	Priority          int            `json:"priority" comment:"-1~2"`
	Style             int            `json:"style,omitempty"`
	AlertType         int            `json:"alert_type,omitempty"`
	BigText           string         `json:"big_text,omitempty"`
	LargeIcon         string         `json:"large_icon,omitempty"`
	Inbox             any            `json:"inbox,omitempty"`
	BigPicPath        string         `json:"big_pic_path,omitempty"`
	SmallIconUri      string         `json:"small_icon_uri,omitempty"`
	UriActivity       string         `json:"uri_activity,omitempty"`
	UriAction         string         `json:"uri_action,omitempty"`
	BadgeAddNum       int            `json:"badge_add_num,omitempty"`
	BadgeSetNum       int            `json:"badge_set_num,omitempty"`
	BadgeClass        string         `json:"badge_class,omitempty"`
	Sound             string         `json:"sound,omitempty"`
	ShowBeginTime     string         `json:"show_begin_time,omitempty"`
	ShowEndTime       string         `json:"show_end_time,omitempty"`
	DisplayForeground string         `json:"display_foreground,omitempty"`
	Intent            Intent         `json:"intent,omitempty"`
	Extras            map[string]any `json:"extras,omitempty"`
}

type Ios struct {
	Alert             string         `json:"alert"`
	Sound             string         `json:"sound,omitempty"`
	Badge             string         `json:"badge,omitempty"`
	ContentAvailable  bool           `json:"content-available,omitempty"`
	MutableContent    bool           `json:"mutable-content,omitempty"`
	Category          string         `json:"category,omitempty"`
	ThreadId          string         `json:"thread-id,omitempty"`
	Extras            map[string]any `json:"extras,omitempty"`
	InterruptionLevel string         `json:"interruption-level,omitempty"`
}

type SmsMessage struct {
	TempId   int `json:"temp_id"`
	TempPara struct {
		Code string `json:"code"`
	} `json:"temp_para"`
	DelayTime    int  `json:"delay_time"`
	ActiveFilter bool `json:"active_filter"`
}

type Options struct {
	SendNo            int    `json:"sendno,omitempty"`
	TimeToLive        int    `json:"time_to_live,omitempty"`
	OverrideMsgID     int64  `json:"override_msg_id,omitempty"`
	ApnsProduction    bool   `json:"apns_production,omitempty"`
	ApnsCollapseId    string `json:"apns_collapse_id,omitempty"`
	BigPushDuration   int    `json:"big_push_duration,omitempty"`
	ThirdPartyChannel any    `json:"third_party_channel,omitempty"`
	Classification    int    `json:"classification,omitempty"`
	TargetEvent       any    `json:"target_event,omitempty"`
}

type Callback struct {
	Url    string `json:"url"`
	Params any    `json:"params"`
	Type   int    `json:"type"`
}

type Intent struct {
	Url string `json:"url"`
}

type Notification3rd struct {
	Title       string         `json:"title,omitempty"`
	Content     string         `json:"content"`
	Intent      Intent         `json:"intent,omitempty"`
	UriActivity string         `json:"uri_activity,omitempty"`
	UriAction   string         `json:"uri_action,omitempty"`
	BadgeAddNum int            `json:"badge_add_num,omitempty"`
	BadgeSetNum int            `json:"badge_set_num,omitempty"`
	BadgeClass  string         `json:"badge_class,omitempty"`
	Sound       string         `json:"sound,omitempty"`
	ChannelID   string         `json:"channel_id,omitempty"`
	Extras      map[string]any `json:"extras,omitempty"`
}
