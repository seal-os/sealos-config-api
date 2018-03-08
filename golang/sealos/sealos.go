/*
   Copyright (c) 2018 Open Devices. All rights reserved.
*/

package sealos

type APISealOSConfig struct {
        Api_User_Name           string  `json:"API_USER_NAME"`
        Api_Project_Name        string  `json:"API_PROJECT_NAME"`
        Api_Key                 string  `json:"API_KEY"`
        Api_Key_Subscribe       string  `json:"API_KEY_SUBSCRIBE"`
        Api_Key_Publish         string  `json:"API_KEY_PUBLISH"`


        System_Hostname         string  `json:"SYSTEM_HOSTNAME"`
        System_Dns_Servers      []string        `json:"SYSTEM_DNS_SERVERS"`
        System_Ntp_Servers      []string        `json:"SYSTEM_NTP_SERVERS"`
        System_Manager_Port     int     `json:"SYSTEM_MANAGER_PORT"`
        System_Log_Storage      string  `json:"SYSTEM_LOG_STORAGE"`

        System_Log_Level        string  `json:"SYSTEM_LOG_LEVEL"`
        System_Log_Max          string  `json:"SYSTEM_LOG_MAX"`
        System_Log_Compress     string  `json:"SYSTEM_LOG_COMPRESS"`
        System_Log_Seal         string  `json:"SYSTEM_LOG_SEAL"`

        Runtime_Watchdog_Sec    string  `json:"RUNTIME_WATCHDOG_SEC"`
        Shutdown_Watchdog_Sec   string  `json:"SHUTDOWN_WATCHDOG_SEC"`
        Watchdog_Device         string  `json:"WATCHDOG_DEVICE"`


        /* Internal configuration */
        Api_Endpoint            string  `json:"API_ENDPOINT"`
        Api_User_Id             string  `json:"API_USER_ID"`
        Api_Project_Id          string  `json:"API_PROJECT_ID"`
        Api_Project_Device_Type string  `json:"API_PROJECT_DEVICE_TYPE"`

        /* This Should be the same as /etc/machine-id file */
        Api_Device_UUID         string  `json:"API_DEVICE_UUID"`
}
