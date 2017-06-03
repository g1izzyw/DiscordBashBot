package configuration

type BotConfigurationObject struct {
	Token            *string
	ValidChannelList []string
	RedisHost        *string
	RedisPort        *string
	RedisPassword    *string
	RedisUnixSocket  *string
}

var BotConfiguration *BotConfigurationObject
