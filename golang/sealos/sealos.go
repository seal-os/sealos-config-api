/*
   Copyright (c) 2020 Open Devices GmbH. All rights reserved.
*/

package sealos

type APIKeyHeader struct {
	Alg string `json:"Alg,omitempty"`
	Typ string `json:"Typ,omitempty"`
}

type APIKeyPayload struct {
	/* Usage count of this current key */
	Usage_Count int64 `json:"usage_count,omitempty"`

	/* Age when the key/Access was generated */
	Age int64 `json:"age,omitempty"`

	/*
	   Token/Key Use: refresh/access
	*/
	Token_Use string `json:"token_use,omitempty"`

	/* Expiration time of the key */
	Exp int64 `json:"exp,omitempty"`

	/* Subject ID to access */
	Sub string `json:"sub,omitempty"`

	/* Subject ID to access */
	Username string `json:"username,omitempty"`

	/* Subject ID to access */
	Client_id string `json:"client_id,omitempty"`

	/* Access role on the subject is used in backend only */
	Role string `json:"role,omitempty"`

	/* Issuer */
	Iss string `json:"iss,omitempty"`

	/*
	   Ionoid Resource Access
	   res::$ORG_RESOURCE_ID::$PROJECT_ID
	*/
	Ionoid_Res string `json:"ionoid_res,omitempty"`

	/*
	   Ionoid Scope Access
	   scope::$STRING:: $STRING || $ID || $URL
	*/
	Ionoid_Scope string `json:"ionoid_scope,omitempty"`

	/* Audiance: The ID or Project ID where this entity/device is registered */
	Aud string `json:"aud,omitempty"`

	/* True if this key is obsolete */
	Obsoleted bool `json:"Obsoleted,omitempty"`
}

type APIKey struct {
	Header    APIKeyHeader  `json:"Header,omitempty"`
	Payload   APIKeyPayload `json:"Payload,omitempty"`
	Signature string        `json:"Signature,omitempty"`
	Key       string        `json:"Key,omitempty"`
}

/* WIFI API do not use omitempty here */
type APIWifi struct {
	SSID           string `json:"SSID,omitempty"`
	BSS            string `json:"BSS,omitempty"`
	Interface      string `json:"INTERFACE,omitempty"`
	Scan_SSID      int    `json:"SCAN_SSID,omitempty"`
	Password       string `json:"PASSWORD,omitempty"`
	Psk            string `json:"PSK,omitempty"`
	Security       string `json:"SECURITY,omitempty"`
	Priority       int    `json:"PRIORITY,omitempty"`
	Wep_TX_Keyidx  int    `json:"WEP_TX_KEYIDX,omitempty"`
	Wep_Key0       string `json:"WEP_KEY0,omitempty"`
	Fallback       string `json:"FALLBACK,omitempty"`
	Associated     string `json:"ASSOCIATED,omitempty"`
	Signal_Quality uint   `json:"SIGNAL_QUALITY,omitempty"`
	RxRateS        string `json:"RXRATE_S,omitempty"`
	TxRateS        string `json:"TXRATE_S,omitempty"`
}

type APICellular struct {
	/* Network ID as set in the config file */
	NetworkId string `json:"NETWORK_ID,omitempty"`
	/* Interface to be used */
	Interface     string   `json:"INTERFACE,omitempty"`
	Autoconfig    bool     `json:"AUTOCONFIG,omitempty"`
	Autoconnect   bool     `json:"AUTOCONNECT,omitempty"`
	Priority      int      `json:"PRIORITY,omitempty"`
	Apn           string   `json:"APN,omitempty"`
	Plan          string   `json:"PLAN,omitempty"`
	Usage         string   `json:"USAGE,omitempty"`
	Username      string   `json:"USERNAME,omitempty"`
	Password      string   `json:"PASSWORD,omitempty"`
	Pin           string   `json:"PIN,omitempty"`
	SimId         string   `json:"SIM_ID,omitempty"`
	SimOperatorId string   `json:"SIM_OPERATOR_ID,omitempty"`
	Dns           []string `json:"DNS,omitempty"`
}

type Connectivity_Check struct {
	Uri string `json:"URI,omitempty"`

	/*
	   Connectivity check interval in seconds, if less or equal 0
	   means connectivity check is disabled.
	   Possible values are: 0, 60, 300, 600, 1200

	   By default 1200seconds.
	*/
	Interval int `json:"INTERVAL,omitempty"`

	/*
		           The number of failed connectivity checks required to reconfigure
		           the network including modems.

		           If less or qual 0 means reconfigure is disabled.

		           By default 0. Valid values: 0, 1, 2, 4, 8

			   Reconfiguring network will run the commands in `MODEM_RESET`
			   and will stop both Network Manager and Modem Manager then it
			   will restart them again.
	*/
	Reconfigure_after int `json:"RECONFIGURE_AFTER,omitempty"`

	/*
	   Modem reset instructions, possible values separated by ',':
	   disable:                             Disable a given modem
	   set-power-state-low/power-low        Set low power state in the modem
	   set-power-state-off/power-off:       Power off the modem
	   reset:                               Reset the modems

	   Example of multiple combination:
	        disable,power-off,reset
	*/
	ModemReset string `json:"MODEM_RESET,omitempty"`

	/*
	   Number of maximum reconfigure retries before failing definitly
	   and execute what the Failure_Action points to.

	   If less or equal 0 means endless, no limit
	   0 means no limit

	   Possible values: 0, 1, 2, 4, 8
	*/
	Reconfigure_limit int `json:"RECONFIGURE_LIMIT,omitempty"`

	/*
			   Action to take when we reach maximum reconfiguration limit:

		           Which is: interval * (reconfigure_after + 1) in time seconds/minutes

			   Possible values are:
		                none    (No action) Do not take any action and service state may stay in failure. Default.
			        reboot  (Normal reboot) Performs a normal reboot operation.
			        reboot-force (Force reboot) Performs a forced reboot operation by terminating all processes, should cause no dirty file systems.
			        reboot-immediate (Immediate reboot) Similar to a hardware reboot, which might result in data loss.
			        poweroff (Normal poweroff) Performs a normal poweroff operation.
			        poweroff-force (Force poweroff) Performs a forced poweroff operation by terminating all processes.
			        poweroff-immediate (Immediate poweroff) Similar to a hardware poweroff operation, which might result in data loss.


		           For later:
			        exit
			        exit-force


	*/
	Failure_Action string `json:"FAILURE_ACTION,omitempty"`
}

type APISystemNetwork struct {
	WifiCountry       string             `json:"WIFI_COUNTRY,omitempty"`
	MacAddrProtection string             `json:"MAC_ADDR_PROTECTION,omitempty"`
	Wifi              APIWifi            `json:"WIFI,omitempty"`
	WifiNetworks      []APIWifi          `json:"WIFI_NETWORKS,omitempty"`
	CellularNetworks  []APICellular      `json:"CELLULAR_NETWORKS,omitempty"`
	ConnectivityCheck Connectivity_Check `json:"CONNECTIVITY_CHECK,omitempty"`
}

/* System and Apps deployment settings */
type APIDeploymentSettings struct {
	App_Dual_AB       string `json:"APP_DUAL_AB,omitempty"`
	App_Delta_Updates string `json:"APP_DELTA_UPDATES,omitempty"`
}

/* SealOS Config .json file do not use omitempty here */
type APISealOSConfig struct {
	Api_Unmanaged bool `json:"ApiUnmanaged"`

	Api_Project_Name string `json:"API_PROJECT_NAME"`
	Api_Device_Name  string `json:"API_DEVICE_NAME"`
	Api_Key_Devices  string `json:"API_KEY_DEVICES"`

	/* Api Endpoint Keys/Endpoints */
	Api_Auth_Keys []APIKey `json:"ApiAuthKeys,omitempty"`

	/* TODO: TO BE DELETED */
	Api_Key_Publish_Events   string `json:"API_KEY_PUBLISH_EVENTS"`
	Api_Key_Subscribe_Events string `json:"API_KEY_SUBSCRIBE_EVENTS"`

	Api_Key_Publish_Responses   string `json:"API_KEY_PUBLISH_RESPONSES"`
	Api_Key_Subscribe_Responses string `json:"API_KEY_SUBSCRIBE_RESPONSES"`

	/* TODO: END OF TO BE DELETED */

	System_Dns_Servers []string         `json:"SYSTEM_DNS_SERVERS"`
	System_Ntp_Servers []string         `json:"SYSTEM_NTP_SERVERS"`
	System_Network     APISystemNetwork `json:"SYSTEM_NETWORK"`

	Deployment_Settings APIDeploymentSettings `json:"DEPLOYMENT_SETTINGS"`

	System_Manager_Port int    `json:"SYSTEM_MANAGER_PORT"`
	System_Log_Storage  string `json:"SYSTEM_LOG_STORAGE"`

	System_Log_Level    string `json:"SYSTEM_LOG_LEVEL"`
	System_Log_Max      string `json:"SYSTEM_LOG_MAX"`
	System_Log_Compress string `json:"SYSTEM_LOG_COMPRESS"`
	System_Log_Seal     string `json:"SYSTEM_LOG_SEAL"`

	Kernel_Boot_Args string `json:"KERNEL_BOOT_ARGS"`

	/* Internal configuration */
	Api_Endpoint_Devices string `json:"API_ENDPOINT_DEVICES"`

	Api_Resource_Org_Id    string `json:"API_RESOURCE_ORG_ID"`
	Api_Project_Id         string `json:"API_PROJECT_ID"`
	Api_Project_M2M_Id     string `json:"API_PROJECT_M2M_ID,omitempty"`
	Api_Project_Network_Id string `json:"API_PROJECT_NETWORK_ID,omitempty"`

	Api_User_Id string `json:"API_USER_ID"`

	Api_Project_Device_Arch string `json:"API_PROJECT_DEVICE_ARCH"`
	Api_Project_OS          string `json:"API_PROJECT_OS"`

	Api_Device_Vendor string `json:"API_DEVICE_VENDOR"`

	Api_Device_OS_Release string `json:"API_DEVICE_OS_RELEASE"`

	/* Device Platform Model */
	Api_Device_Model       string `json:"API_DEVICE_MODEL"`
	Api_Device_Short_Model string `json:"API_DEVICE_SHORT_MODEL"`

	/* This Should be the same as /etc/machine-id file */
	Api_Device_UUID string `json:"API_DEVICE_UUID"`

	/* This is a local UUID auto generated and private to IONOID */
	Api_Device_Static_UUID string `json:"API_DEVICE_STATIC_UUID"`

	/*
	 * This is a Network UUID derived from Device UUID and can be
	 * used in untrusted networks to connects to services on
	 * Devices.
	 */
	Api_Device_Network_UUID string `json:"API_DEVICE_NETWORK_UUID"`

	/*
	 * This is the M2M Client UUID derived from Device UUID and can be
	 * used to construct M2M HTTP MQTT TOPICS, channels or in any
	 * Publish/Subscribe Mechanism
	 */
	Api_Device_M2M_UUID string `json:"API_DEVICE_M2M_UUID"`

	/*
	 * This is the Public UUID that can be used inside messages and
	 * shared with untrusted parties. Used to identify messages
	 * originated from a specific Device
	 */
	Api_Device_Public_UUID string `json:"API_DEVICE_PUBLIC_UUID"`

	/*
	   Auto Send Apps Logs when they start: true or false
	*/
	Api_Device_Apps_Auto_Log string `json:"API_DEVICE_APPS_AUTO_LOG"`

	Api_Device_Production bool `json:"API_DEVICE_PRODUCTION"`

	Api_Project_Device_Tags map[string]string `json:"API_PROJECT_DEVICE_TAGS,omitempty"`
	Api_Project_Env_Vars    map[string]string `json:"API_PROJECT_ENV_VARS,omitempty"`
	Api_Device_Env_Vars     map[string]string `json:"API_DEVICE_ENV_VARS,omitempty"`

	Api_Device_Manager_Version string `json:"API_DEVICE_MANAGER_VERSION"`
}

type APISealOSActionRequest struct {
	ID      string `json:"id,omitempty"`
	Action  string `json:"action"`
	Command string `json:"command,omitempty"`
}

type APISealOSStatusRequest struct {
	Api_Device_UUID string `json:"API_DEVICE_UUID,omitempty"`
	Action          string `json:"action"`
}

type APISealOSStatusResponse struct {
	Api_Device_UUID   string `json:"API_DEVICE_UUID"`
	Api_Device_Status string `json:"API_DEVICE_STATUS"`
}

type APISealOSLogsRequest struct {
	Api_Device_UUID string      `json:"API_DEVICE_UUID"`
	Action          string      `json:"action"`
	Command         string      `json:"command,omitempty"`
	Format          string      `json:"format,omitempty"`
	Publisher       interface{} `json:"publisher,omitempty"`
}

type APISealOSLogsResponse struct {
	Api_Device_UUID string `json:"API_DEVICE_UUID"`
	Format          string `json:"format,omitempty"`
	Logs            []byte `json:"logs,omitempty"`
}
