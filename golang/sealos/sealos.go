/*
   Copyright (c) 2018 Open Devices. All rights reserved.
*/

package sealos

const (
        SealOSManagerPath       = "/usr/lib/sealos-manager/"
        SealOSFactoryPath       = "/usr/lib/sealos-manager/factory/"
)

type APIWifi struct {
        SSID            string                  `json:"SSID"`
        Password        string                  `json:"PASSWORD"`
        Security        string                  `json:"SECURITY"`
}

type APISystemNetwork struct {
        Wifi            APIWifi                 `json:"WIFI"`
}

type APISealOSConfig struct {
        Api_User_Email          string  `json:"API_USER_EMAIL"`
        Api_Project_Name        string  `json:"API_PROJECT_NAME"`
        Api_Device_Name         string  `json:"API_DEVICE_NAME"`
        Api_Key_Devices         string  `json:"API_KEY_DEVICES"`
        Api_Key_App             string  `json:"API_KEY_APP"`
        Api_Key_Subscribe_Events        string  `json:"API_KEY_SUBSCRIBE_EVENTS"`
        Api_Key_Publish_Responses       string  `json:"API_KEY_PUBLISH_RESPONSES"`

        System_Hostname         string  `json:"SYSTEM_HOSTNAME"`
        System_Dns_Servers      []string        `json:"SYSTEM_DNS_SERVERS"`
        System_Ntp_Servers      []string        `json:"SYSTEM_NTP_SERVERS"`
        System_Network          APISystemNetwork        `json:"SYSTEM_NETWORK"`

        System_Manager_Port     int     `json:"SYSTEM_MANAGER_PORT"`
        System_Log_Storage      string  `json:"SYSTEM_LOG_STORAGE"`

        System_Log_Level        string  `json:"SYSTEM_LOG_LEVEL"`
        System_Log_Max          string  `json:"SYSTEM_LOG_MAX"`
        System_Log_Compress     string  `json:"SYSTEM_LOG_COMPRESS"`
        System_Log_Seal         string  `json:"SYSTEM_LOG_SEAL"`

        Kernel_Boot_Args        string  `json:"KERNEL_BOOT_ARGS"`

        Runtime_Watchdog_Sec    string  `json:"RUNTIME_WATCHDOG_SEC"`
        Shutdown_Watchdog_Sec   string  `json:"SHUTDOWN_WATCHDOG_SEC"`
        Watchdog_Device         string  `json:"WATCHDOG_DEVICE"`


        /* Internal configuration */
        Api_Endpoint_Devices    string  `json:"API_ENDPOINT_DEVICES"`
        Api_User_Id             string  `json:"API_USER_ID"`
        Api_Project_Id          string  `json:"API_PROJECT_ID"`
        Api_Project_Device_Arch string  `json:"API_PROJECT_DEVICE_ARCH"`
        Api_Project_OS          string  `json:"API_PROJECT_OS"`

        Api_Device_Vendor       string  `json:"API_DEVICE_VENDOR"`

        Api_Device_OS_Release   string  `json:"API_DEVICE_OS_RELEASE"`

        /* This Should be the same as /etc/machine-id file */
        Api_Device_UUID         string  `json:"API_DEVICE_UUID"`

        /* This is a local UUID auto generated and private */
        Api_Device_Static_UUID  string  `json:"API_DEVICE_STATIC_UUID"`

        /*
         * This is a Network UUID derived from Device UUID and can be
         * used in untrusted networks
         */
        Api_Device_Network_UUID string  `json:"API_DEVICE_NETWORK_UUID"`

        /*
         * This is the MQTT Client UUID derived from Device UUID and can be
         * used to construct MQTT TOPICS, channels or in any
         * Publish/Subscribe Mechanism
         */
        Api_Device_Mqtt_UUID    string  `json:"API_DEVICE_MQTT_UUID"`

        Api_Device_Production   bool    `json:"API_DEVICE_PRODUCTION"`
}

type APISealOSActionRequest struct {
        ID                      string          `json:"id,omitempty"`
        Action                  string          `json:"action"`
        Command                 string          `json:"command,omitempty"`
}

type APISealOSStatusRequest struct {
        Api_Device_UUID         string          `json:"API_DEVICE_UUID,omitempty"`
        Action                  string          `json:"action"`
}

type APISealOSStatusResponse struct {
        Api_Device_UUID         string          `json:"API_DEVICE_UUID"`
        Api_Device_Status       string          `json:"API_DEVICE_STATUS"`
}

type APISealOSLogsRequest struct {
        Api_Device_UUID         string          `json:"API_DEVICE_UUID"`
        Action                  string          `json:"action"`
        Command                 string          `json:"command,omitempty"`
        Format                  string          `json:"format,omitempty"`
        Publisher               interface{}     `json:"publisher,omitempty"`
}

type APISealOSLogsResponse struct {
        Api_Device_UUID         string          `json:"API_DEVICE_UUID"`
        Format                  string          `json:"format,omitempty"`
        Logs                    []byte          `json:"logs,omitempty"`
}
