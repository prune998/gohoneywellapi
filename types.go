type TSerie []struct {
	LocationID int    `json:"locationID"`
	Name       string `json:"name"`
	Country    string `json:"country"`
	Zipcode    string `json:"zipcode"`
	Devices		 []Device `json:"Devices"`
}

type Device    struct {
		DeviceDetails struct {
			Configurations struct {
				Model           string `json:"model"`
				FirmwareVersion string `json:"firmwareVersion"`
				CommAPIVersion  string `json:"commApiVersion"`
				Mac             string `json:"mac"`
				SerialNumber    string `json:"serialNumber"`
				WifiInfo        struct {
					Name string `json:"name"`
				} `json:"wifiInfo"`
				FirmwareUpgradeInfo struct {
					LastSuccessfulUpgradeTime     time.Time `json:"lastSuccessfulUpgradeTime"`
					LastUpgradeStartTime          time.Time `json:"lastUpgradeStartTime"`
					ExpectedUpgradeCompletionTime time.Time `json:"expectedUpgradeCompletionTime"`
				} `json:"firmwareUpgradeInfo"`
				PrimaryZWaveController struct {
					ID             string `json:"id"`
					Name           string `json:"name"`
					Configurations struct {
						Version                 string `json:"version"`
						HomeID                  string `json:"homeId"`
						Type                    string `json:"type"`
						NodeID                  string `json:"nodeId"`
						NodeDescription         string `json:"nodeDescription"`
						AutomationDeviceType    string `json:"automationDeviceType"`
						AutomationDeviceSubType string `json:"automationDeviceSubType"`
						ManufacturerName        string `json:"manufacturerName"`
						ManufacturerID          string `json:"manufacturerId"`
						ProductID               string `json:"productId"`
						ProductType             string `json:"productType"`
						RouterCapability        bool   `json:"routerCapability"`
						AssociationState        bool   `json:"associationState"`
						DeviceSecurity          bool   `json:"deviceSecurity"`
						SpecificType            int    `json:"specificType"`
						SecurityMode            string `json:"securityMode"`
					} `json:"configurations"`
					State struct {
						Value         int  `json:"value"`
						IsConnected   bool `json:"isConnected"`
						InReplaceMode bool `json:"inReplaceMode"`
					} `json:"state"`
				} `json:"primaryZWaveController"`
				VoiceCommands []struct {
					VoiceEngine string `json:"voiceEngine"`
					Enabled     bool   `json:"enabled"`
				} `json:"voiceCommands"`
			} `json:"configurations"`
			State struct {
				IsACLoss     bool `json:"isACLoss"`
				BatteryState struct {
					IsLow bool `json:"isLow"`
					Level int  `json:"level"`
				} `json:"batteryState"`
				IsTampered                         bool      `json:"isTampered"`
				IsOnline                           bool      `json:"isOnline"`
				DiscoveryMode                      string    `json:"discoveryMode"`
				LastCommunicationStatusChange      time.Time `json:"lastCommunicationStatusChange"`
				IsSyncInProgress                   bool      `json:"isSyncInProgress"`
				IsUpgradeInProgress                bool      `json:"isUpgradeInProgress"`
				AlexaConnectionState               string    `json:"alexaConnectionState"`
				IsMandatoryFirmwareUpgradeRequired bool      `json:"isMandatoryFirmwareUpgradeRequired"`
			} `json:"state"`
			Partitions []struct {
				ID             string `json:"id"`
				Name           string `json:"name"`
				Configurations struct {
					CanRestartTimer  bool `json:"canRestartTimer"`
					EntryDelay       int  `json:"entryDelay"`
					ExitDelay        int  `json:"exitDelay"`
					EnableDeterrence bool `json:"enableDeterrence"`
					SetStateConfigs  []struct {
						SetState         string `json:"setState"`
						EnableAlarm      bool   `json:"enableAlarm"`
						EnableCamera     bool   `json:"enableCamera"`
						ChimeVolume      int    `json:"chimeVolume"`
						EnableDeterrence bool   `json:"enableDeterrence"`
						EnableOsmv       bool   `json:"enableOsmv"`
					} `json:"setStateConfigs"`
					DeterrenceActions struct {
						Chimes []string `json:"chimes"`
					} `json:"deterrenceActions"`
				} `json:"configurations"`
				State struct {
					SetState                string    `json:"setState"`
					InEntryDelay            bool      `json:"inEntryDelay"`
					InExitDelay             bool      `json:"inExitDelay"`
					AlarmState              string    `json:"alarmState"`
					IsTrouble               bool      `json:"isTrouble"`
					IsFault                 bool      `json:"isFault"`
					IsBypassed              bool      `json:"isBypassed"`
					LastUpdatedTime         time.Time `json:"lastUpdatedTime"`
					SilentModeEnabled       bool      `json:"silentModeEnabled"`
					LastSetStateUpdatedTime time.Time `json:"lastSetStateUpdatedTime"`
				} `json:"state"`
			} `json:"partitions"`
			Sensors []struct {
				ID             string `json:"id"`
				Name           string `json:"name"`
				Configurations struct {
					FirmwareVersion string `json:"firmwareVersion"`
					SerialNumber    string `json:"serialNumber"`
					BridgeID        string `json:"bridgeId"`
					Sensitivity     int    `json:"sensitivity"`
					Chime           string `json:"chime"`
					IsEnabled       bool   `json:"isEnabled"`
					SensorCommType  string `json:"sensorCommType"`
					ResponseType    string `json:"responseType"`
					SensorType      string `json:"sensorType"`
					SensorSubType   string `json:"sensorSubType"`
				} `json:"configurations"`
				State struct {
					IsAlarm           bool `json:"isAlarm"`
					IsBypassed        bool `json:"isBypassed"`
					IsFault           bool `json:"isFault"`
					IsTrouble         bool `json:"isTrouble"`
					IsTampered        bool `json:"isTampered"`
					SupervisionFailed bool `json:"supervisionFailed"`
					IsEnrolled        bool `json:"isEnrolled"`
					IsConnected       bool `json:"isConnected"`
					SignalStrength    int  `json:"signalStrength"`
					BatteryState      struct {
						IsLow bool `json:"isLow"`
						Level int  `json:"level"`
					} `json:"batteryState"`
					ActiveTests []interface{} `json:"activeTests"`
					IsVerified  bool          `json:"isVerified"`
				} `json:"state"`
			} `json:"sensors"`
			KeyFobs []struct {
				ID             string `json:"id"`
				Name           string `json:"name"`
				PartitionID    string `json:"partitionId"`
				Configurations struct {
					SerialNumber    string `json:"serialNumber"`
					Size            int    `json:"size"`
					FirmwareVersion string `json:"firmwareVersion"`
				} `json:"configurations"`
				State struct {
					KeyCommandMapping []struct {
						KeyIndex     int `json:"keyIndex"`
						CommandIndex int `json:"commandIndex"`
					} `json:"keyCommandMapping"`
					IsEnrolled   bool `json:"isEnrolled"`
					BatteryState struct {
						IsLow bool `json:"isLow"`
						Level int  `json:"level"`
					} `json:"batteryState"`
					SignalStrength int `json:"signalStrength"`
				} `json:"state"`
			} `json:"keyFobs"`
			AutomationDevices struct {
				Switches []struct {
					Configurations struct {
						SwitchType              string `json:"switchType"`
						SwitchSubType           string `json:"switchSubType"`
						NodeID                  string `json:"nodeId"`
						NodeDescription         string `json:"nodeDescription"`
						AutomationDeviceType    string `json:"automationDeviceType"`
						AutomationDeviceSubType string `json:"automationDeviceSubType"`
						ManufacturerName        string `json:"manufacturerName"`
						ManufacturerID          string `json:"manufacturerId"`
						ProductID               string `json:"productId"`
						ProductType             string `json:"productType"`
						RouterCapability        bool   `json:"routerCapability"`
						AssociationState        bool   `json:"associationState"`
						DeviceSecurity          bool   `json:"deviceSecurity"`
						SpecificType            int    `json:"specificType"`
						SecurityMode            string `json:"securityMode"`
					} `json:"configurations"`
					State struct {
						LastCommunicationStatusChangeTime time.Time `json:"lastCommunicationStatusChangeTime"`
						Value                             int       `json:"value"`
						IsConnected                       bool      `json:"isConnected"`
						InReplaceMode                     bool      `json:"inReplaceMode"`
					} `json:"state"`
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"switches"`
			} `json:"automationDevices"`
			OnboardDevices []struct {
				DeviceClass           string `json:"deviceClass"`
				DeviceType            string `json:"deviceType"`
				DeviceID              string `json:"deviceID"`
				MaxResolution         string `json:"maxResolution"`
				Manufacturer          string `json:"manufacturer"`
				Model                 string `json:"model"`
				Serial                string `json:"serial"`
				FirmwareVer           string `json:"firmwareVer"`
				MacID                 string `json:"macID"`
				UserDefinedDeviceName string `json:"userDefinedDeviceName"`
				IsAlive               bool   `json:"isAlive"`
				IsUpgrading           bool   `json:"isUpgrading"`
				IsProvisioned         bool   `json:"isProvisioned"`
				WifiStrength          string `json:"wifiStrength"`
				Backend               string `json:"backend"`
				PeopleDetectionState  string `json:"peopleDetectionState"`
			} `json:"onboardDevices"`
		} `json:"deviceDetails,omitempty"`
		GeoFence        string `json:"geoFence,omitempty"`
		FrConfiguration struct {
			MaxPersons int    `json:"maxPersons"`
			Plan       string `json:"plan"`
		} `json:"frConfiguration,omitempty"`
		DeviceClass           string `json:"deviceClass"`
		DeviceType            string `json:"deviceType"`
		DeviceID              string `json:"deviceID"`
		UserDefinedDeviceName string `json:"userDefinedDeviceName"`
		Name                  string `json:"name,omitempty"`
		IsAlive               bool   `json:"isAlive"`
		IsProvisioned         bool   `json:"isProvisioned"`
		MacID                 string `json:"macID,omitempty"`
		DeviceSettings        struct {
		} `json:"deviceSettings,omitempty"`
		Service struct {
			Mode string `json:"mode"`
		} `json:"service"`
		FirmwareVersion          string `json:"firmwareVersion,omitempty"`
		DisplayedOutdoorHumidity int    `json:"displayedOutdoorHumidity,omitempty"`
		VacationHold             struct {
			VacationStart string `json:"vacationStart"`
			VacationEnd   string `json:"vacationEnd"`
			HeatSetpoint  int    `json:"heatSetpoint"`
			CoolSetpoint  int    `json:"coolSetpoint"`
			Enabled       bool   `json:"enabled"`
		} `json:"vacationHold,omitempty"`
		CurrentSchedulePeriod struct {
			Day    string `json:"day"`
			Period string `json:"period"`
		} `json:"currentSchedulePeriod,omitempty"`
		ScheduleCapabilities struct {
			AvailableScheduleTypes []string `json:"availableScheduleTypes"`
			SchedulableFan         bool     `json:"schedulableFan"`
		} `json:"scheduleCapabilities,omitempty"`
		ScheduleType struct {
			ScheduleType    string `json:"scheduleType"`
			ScheduleSubType string `json:"scheduleSubType"`
		} `json:"scheduleType,omitempty"`
		ScheduleStatus        string `json:"scheduleStatus,omitempty"`
		AllowedTimeIncrements int    `json:"allowedTimeIncrements,omitempty"`
		Settings              struct {
			HardwareSettings struct {
				Brightness    int `json:"brightness"`
				MaxBrightness int `json:"maxBrightness"`
			} `json:"hardwareSettings"`
			Fan struct {
				AllowedModes     []string `json:"allowedModes"`
				ChangeableValues struct {
					Mode string `json:"mode"`
				} `json:"changeableValues"`
				AllowedSpeeds []struct {
					Item  string `json:"item"`
					Value struct {
						Speed int    `json:"speed"`
						Mode  string `json:"mode"`
					} `json:"value"`
				} `json:"allowedSpeeds"`
			} `json:"fan"`
			TemperatureMode struct {
				Air bool `json:"air"`
			} `json:"temperatureMode"`
			SpecialMode struct {
			} `json:"specialMode"`
		} `json:"settings,omitempty"`
		IsUpgrading    bool `json:"isUpgrading,omitempty"`
		ParentDeviceID int  `json:"parentDeviceId,omitempty"`
		PartnerInfo    struct {
			SingleOrMultiODUConfiguration int    `json:"singleOrMultiODUConfiguration"`
			ParentDeviceModelID           int    `json:"parentDeviceModelId"`
			ParentDeviceBrandID           int    `json:"parentDeviceBrandId"`
			OduName                       string `json:"oduName"`
		} `json:"partnerInfo,omitempty"`
		DeviceRegistrationDate string   `json:"deviceRegistrationDate,omitempty"`
		Units                  string   `json:"units,omitempty"`
		IndoorTemperature      int      `json:"indoorTemperature,omitempty"`
		OutdoorTemperature     float64  `json:"outdoorTemperature,omitempty"`
		AllowedModes           []string `json:"allowedModes,omitempty"`
		Deadband               int      `json:"deadband,omitempty"`
		HasDualSetpointStatus  bool     `json:"hasDualSetpointStatus,omitempty"`
		MinHeatSetpoint        int      `json:"minHeatSetpoint,omitempty"`
		MaxHeatSetpoint        int      `json:"maxHeatSetpoint,omitempty"`
		MinCoolSetpoint        int      `json:"minCoolSetpoint,omitempty"`
		MaxCoolSetpoint        int      `json:"maxCoolSetpoint,omitempty"`
		ChangeableValues       struct {
			Mode                     string `json:"mode"`
			HeatSetpoint             int    `json:"heatSetpoint"`
			CoolSetpoint             int    `json:"coolSetpoint"`
			ThermostatSetpointStatus string `json:"thermostatSetpointStatus"`
			NextPeriodTime           string `json:"nextPeriodTime"`
			EndHeatSetpoint          int    `json:"endHeatSetpoint"`
			EndCoolSetpoint          int    `json:"endCoolSetpoint"`
			HeatCoolMode             string `json:"heatCoolMode"`
		} `json:"changeableValues,omitempty"`
		OperationStatus struct {
			Mode                  string `json:"mode"`
			FanRequest            bool   `json:"fanRequest"`
			CirculationFanRequest bool   `json:"circulationFanRequest"`
		} `json:"operationStatus,omitempty"`
		IndoorHumidity        int    `json:"indoorHumidity,omitempty"`
		IndoorHumidityStatus  string `json:"indoorHumidityStatus,omitempty"`
		DeviceModel           string `json:"deviceModel,omitempty"`
		WaterPresent          bool   `json:"waterPresent,omitempty"`
		CurrentSensorReadings struct {
			Time        string  `json:"time"`
			Temperature float64 `json:"temperature"`
			Humidity    float64 `json:"humidity"`
		} `json:"currentSensorReadings,omitempty"`
		CurrentAlarms              []interface{} `json:"currentAlarms,omitempty"`
		LastCheckin                string        `json:"lastCheckin,omitempty"`
		LastDeviceSettingUpdatedOn string        `json:"lastDeviceSettingUpdatedOn,omitempty"`
		BatteryRemaining           int           `json:"batteryRemaining,omitempty"`
		IsRegistered               bool          `json:"isRegistered,omitempty"`
		HasDeviceCheckedIn         bool          `json:"hasDeviceCheckedIn,omitempty"`
		IsDeviceOffline            bool          `json:"isDeviceOffline,omitempty"`
		FirstFailedAttemptTime     string        `json:"firstFailedAttemptTime,omitempty"`
		FailedConnectionAttempts   int           `json:"failedConnectionAttempts,omitempty"`
		WifiSignalStrength         int           `json:"wifiSignalStrength,omitempty"`
		Time                       string        `json:"time,omitempty"`
		Backend                    struct {
		} `json:"backend,omitempty"`
		DeviceSettings struct {
			Temp struct {
				High struct {
					Limit float64 `json:"limit"`
				} `json:"high"`
				Low struct {
					Limit float64 `json:"limit"`
				} `json:"low"`
			} `json:"temp"`
			Humidity struct {
				High struct {
					Limit int `json:"limit"`
				} `json:"high"`
				Low struct {
					Limit int `json:"limit"`
				} `json:"low"`
			} `json:"humidity"`
			UserDefinedName         string `json:"userDefinedName"`
			BuzzerMuted             bool   `json:"buzzerMuted"`
			CheckinPeriod           int    `json:"checkinPeriod"`
			CurrentSensorReadPeriod int    `json:"currentSensorReadPeriod"`
		} `json:"deviceSettings,omitempty"`
		PartnerInfo struct {
			SingleOrMultiODUConfiguration int `json:"singleOrMultiODUConfiguration"`
			ParentDeviceModelID           int `json:"parentDeviceModelId"`
			ParentDeviceBrandID           int `json:"parentDeviceBrandId"`
		} `json:"partnerInfo,omitempty"`
		Groups []struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Rooms []int  `json:"rooms"`
		} `json:"groups,omitempty"`
		VacationHold struct {
			Enabled bool `json:"enabled"`
		} `json:"vacationHold,omitempty"`
		PriorityType     string `json:"priorityType,omitempty"`
		ChangeableValues struct {
			Mode                     string `json:"mode"`
			AutoChangeoverActive     bool   `json:"autoChangeoverActive"`
			HeatSetpoint             int    `json:"heatSetpoint"`
			CoolSetpoint             int    `json:"coolSetpoint"`
			ThermostatSetpointStatus string `json:"thermostatSetpointStatus"`
			NextPeriodTime           string `json:"nextPeriodTime"`
			HeatCoolMode             string `json:"heatCoolMode"`
		} `json:"changeableValues,omitempty"`
		VacationHold struct {
			Enabled bool `json:"enabled"`
		} `json:"vacationHold,omitempty"`
		ChangeableValues struct {
			Mode                     string `json:"mode"`
			AutoChangeoverActive     bool   `json:"autoChangeoverActive"`
			HeatSetpoint             int    `json:"heatSetpoint"`
			CoolSetpoint             int    `json:"coolSetpoint"`
			ThermostatSetpointStatus string `json:"thermostatSetpointStatus"`
			HeatCoolMode             string `json:"heatCoolMode"`
		} `json:"changeableValues,omitempty"`
		VacationHold struct {
			Enabled bool `json:"enabled"`
		} `json:"vacationHold,omitempty"`
		ScheduleType struct {
			ScheduleType string `json:"scheduleType"`
		} `json:"scheduleType,omitempty"`
		ChangeableValues struct {
			Mode                     string `json:"mode"`
			AutoChangeoverActive     bool   `json:"autoChangeoverActive"`
			HeatSetpoint             int    `json:"heatSetpoint"`
			CoolSetpoint             int    `json:"coolSetpoint"`
			ThermostatSetpointStatus string `json:"thermostatSetpointStatus"`
			HeatCoolMode             string `json:"heatCoolMode"`
		} `json:"changeableValues,omitempty"`
		VacationHold struct {
			Enabled bool `json:"enabled"`
		} `json:"vacationHold,omitempty"`
		Settings struct {
			HardwareSettings struct {
				Brightness    int `json:"brightness"`
				MaxBrightness int `json:"maxBrightness"`
			} `json:"hardwareSettings"`
			VentilationModeSettings struct {
				ChangeableValues        string `json:"changeableValues"`
				VentilationTimer        int    `json:"ventilationTimer"`
				VentilationCoreTimer    int    `json:"ventilationCoreTimer"`
				VentilationSpeed        int    `json:"ventilationSpeed"`
				VentilationFanRequested bool   `json:"ventilationFanRequested"`
			} `json:"ventilationModeSettings"`
			Fan struct {
				AllowedModes     []string `json:"allowedModes"`
				ChangeableValues struct {
					Mode string `json:"mode"`
				} `json:"changeableValues"`
			} `json:"fan"`
			TemperatureMode struct {
				Air bool `json:"air"`
			} `json:"temperatureMode"`
			SpecialMode struct {
			} `json:"specialMode"`
		} `json:"settings,omitempty"`
		ChangeableValues struct {
			Mode                     string `json:"mode"`
			HeatSetpoint             int    `json:"heatSetpoint"`
			CoolSetpoint             int    `json:"coolSetpoint"`
			ThermostatSetpointStatus string `json:"thermostatSetpointStatus"`
			HeatCoolMode             string `json:"heatCoolMode"`
		} `json:"changeableValues,omitempty"`
		ThermostatVersion string `json:"thermostatVersion,omitempty"`
		Settings          struct {
			HomeSetPoints struct {
				HomeHeatSP int    `json:"homeHeatSP"`
				HomeCoolSP int    `json:"homeCoolSP"`
				Units      string `json:"units"`
			} `json:"homeSetPoints"`
			AwaySetPoints struct {
				AwayHeatSP   int    `json:"awayHeatSP"`
				AwayCoolSP   int    `json:"awayCoolSP"`
				SmartCoolSP  int    `json:"smartCoolSP"`
				SmartHeatSP  int    `json:"smartHeatSP"`
				UseAutoSmart bool   `json:"useAutoSmart"`
				Units        string `json:"units"`
			} `json:"awaySetPoints"`
			HardwareSettings struct {
				Brightness    int `json:"brightness"`
				Volume        int `json:"volume"`
				MaxBrightness int `json:"maxBrightness"`
				MaxVolume     int `json:"maxVolume"`
			} `json:"hardwareSettings"`
			Fan struct {
				AllowedModes     []string `json:"allowedModes"`
				ChangeableValues struct {
					Mode string `json:"mode"`
				} `json:"changeableValues"`
				FanRunning bool `json:"fanRunning"`
			} `json:"fan"`
			TemperatureMode struct {
				FeelsLike bool `json:"feelsLike"`
				Air       bool `json:"air"`
			} `json:"temperatureMode"`
			SpecialMode struct {
				AutoChangeoverActive bool `json:"autoChangeoverActive"`
				EmergencyHeatActive  bool `json:"emergencyHeatActive"`
			} `json:"specialMode"`
		} `json:"settings,omitempty"`
		Schedule struct {
			ScheduleType string `json:"scheduleType"`
		} `json:"schedule,omitempty"`
		ChangeableValues struct {
			Mode                 string `json:"mode"`
			AutoChangeoverActive bool   `json:"autoChangeoverActive"`
			EmergencyHeatActive  bool   `json:"emergencyHeatActive"`
			HeatSetpoint         int    `json:"heatSetpoint"`
			CoolSetpoint         int    `json:"coolSetpoint"`
			HeatCoolMode         string `json:"heatCoolMode"`
		} `json:"changeableValues,omitempty"`
		OperationStatus struct {
			Mode string `json:"mode"`
		} `json:"operationStatus,omitempty"`
		SmartAway struct {
			Active          bool   `json:"active"`
			TimeOfDay       string `json:"timeOfDay"`
			DurationInHours int    `json:"durationInHours"`
			DurationInDays  int    `json:"durationInDays"`
			LastUsedFormat  string `json:"lastUsedFormat"`
			EndsIn          string `json:"endsIn"`
		} `json:"smartAway,omitempty"`
	} `json:"devices"`
	Users []struct {
		UserID                     int    `json:"userID"`
		Username                   string `json:"username"`
		Firstname                  string `json:"firstname"`
		Lastname                   string `json:"lastname"`
		Created                    int    `json:"created"`
		Deleted                    int64  `json:"deleted"`
		Activated                  bool   `json:"activated"`
		ConnectedHomeAccountExists bool   `json:"connectedHomeAccountExists"`
		LocationRoleMapping        []struct {
			LocationID   int    `json:"locationID"`
			Role         string `json:"role"`
			LocationName string `json:"locationName"`
			Status       int    `json:"status"`
		} `json:"locationRoleMapping"`
		IsOptOut      string `json:"isOptOut"`
		IsCurrentUser bool   `json:"isCurrentUser"`
	} `json:"users"`
	TimeZoneID                string `json:"timeZoneId,omitempty"`
	TimeZone                  string `json:"timeZone"`
	IanaTimeZone              string `json:"ianaTimeZone"`
	DaylightSavingTimeEnabled bool   `json:"daylightSavingTimeEnabled"`
	GeoFences                 []struct {
		GeofenceEnabled bool    `json:"geofenceEnabled"`
		GeoFenceID      int     `json:"geoFenceID"`
		Latitude        float64 `json:"latitude"`
		Longitude       float64 `json:"longitude"`
		Radius          int     `json:"radius"`
		GeoOccupancy    struct {
			WithinFence  int `json:"withinFence"`
			OutsideFence int `json:"outsideFence"`
		} `json:"geoOccupancy"`
	} `json:"geoFences,omitempty"`
	GeoFenceEnabled             bool `json:"geoFenceEnabled"`
	PredictiveAIREnabled        bool `json:"predictiveAIREnabled"`
	ComfortLevel                int  `json:"comfortLevel"`
	GeoFenceNotificationEnabled bool `json:"geoFenceNotificationEnabled"`
	GeoFenceNotificationTypeID  int  `json:"geoFenceNotificationTypeId"`
	Configuration               struct {
		FaceRecognition struct {
			Enabled       bool `json:"enabled"`
			MaxPersons    int  `json:"maxPersons"`
			MaxEtas       int  `json:"maxEtas"`
			MaxEtaPersons int  `json:"maxEtaPersons"`
			Schedules     []struct {
				Time []struct {
					Start string `json:"start"`
					End   string `json:"end"`
				} `json:"time"`
				Days []string `json:"days"`
			} `json:"schedules"`
		} `json:"faceRecognition"`
	} `json:"configuration"`
}