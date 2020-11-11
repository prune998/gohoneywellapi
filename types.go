package hwapi

// App config
type config struct {
	clientCode   string
	token        string
	refreshToken string
}

// REST API conversion to Go
type TSerie struct {
	LocationID                  int        `json:"locationID"`
	Name                        string     `json:"name"`
	Country                     string     `json:"country"`
	Zipcode                     string     `json:"zipcode"`
	Devices                     []Device   `json:"Devices"`
	Users                       []User     `json:"Users"`
	TimeZoneID                  string     `json:"timeZoneId,omitempty"`
	TimeZone                    string     `json:"timeZone"`
	IanaTimeZone                string     `json:"ianaTimeZone"`
	DaylightSavingTimeEnabled   bool       `json:"daylightSavingTimeEnabled"`
	GeoFences                   []GeoFence `json:"geoFences,omitempty"`
	GeoFenceEnabled             bool       `json:"geoFenceEnabled"`
	PredictiveAIREnabled        bool       `json:"predictiveAIREnabled"`
	ComfortLevel                int        `json:"comfortLevel"`
	GeoFenceNotificationEnabled bool       `json:"geoFenceNotificationEnabled"`
	GeoFenceNotificationTypeID  int        `json:"geoFenceNotificationTypeId"`
	Configuration               `json:"configuration"`
}

type Group struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Rooms []int  `json:"rooms"`
}

type Device struct {
	Groups                   []Group `json:"groups"`
	DisplayedOutdoorHumidity int     `json:"displayedOutdoorHumidity,omitempty"`
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

	InBuiltSensorState struct {
		RoomID   int    `json:"roomId"`
		RoomName string `json:"roomName"`
	} `json:"inBuiltSensorState"`

	ScheduleStatus        string `json:"scheduleStatus,omitempty"`
	AllowedTimeIncrements int    `json:"allowedTimeIncrements,omitempty"`

	Settings struct {
		DevicePairingEnabled bool `json:"devicePairingEnabled"`
		HardwareSettings     struct {
			Brightness    int `json:"brightness"`
			MaxBrightness int `json:"maxBrightness"`
		} `json:"hardwareSettings"`
		TemperatureMode struct {
			Air bool `json:"air"`
		} `json:"temperatureMode"`
		SpecialMode struct {
		} `json:"specialMode"`
	}

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
	FirmwareVersion string `json:"firmwareVersion,omitempty"`
	PriorityType    string `json:"priorityType,omitempty"`

	DeviceRegistrationDate string   `json:"deviceRegistrationDate,omitempty"`
	DataSyncStatus         string   `json:"dataSyncStatus,omitempty"`
	Units                  string   `json:"units,omitempty"`
	IndoorTemperature      float64  `json:"indoorTemperature,omitempty"`
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
	IndoorHumidity       int    `json:"indoorHumidity,omitempty"`
	IndoorHumidityStatus string `json:"indoorHumidityStatus,omitempty"`
	DeviceModel          string `json:"deviceModel,omitempty"`
	WaterPresent         bool   `json:"waterPresent,omitempty"`
}

type User struct {
	UserID                     int    `json:"userID"`
	Username                   string `json:"username"`
	Firstname                  string `json:"firstname"`
	Lastname                   string `json:"lastname"`
	Created                    int64  `json:"created"`
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
}

type GeoFence struct {
	GeofenceEnabled bool    `json:"geofenceEnabled"`
	GeoFenceID      int     `json:"geoFenceID"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Radius          int     `json:"radius"`
	GeoOccupancy    struct {
		WithinFence  int `json:"withinFence"`
		OutsideFence int `json:"outsideFence"`
	} `json:"geoOccupancy"`
	GeoFenceNotificationEnabled bool `json:"geoFenceNotificationEnabled"`
}

type Configuration struct {
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
}

// Schedule represent one specific setup in time
type Schedule struct {
	DeviceID        string `json:"deviceID"`
	ScheduleType    string `json:"scheduleType"`
	ScheduleSubType string `json:"scheduleSubType"`

	TimedSchedule struct {
		Days []Day `json:"days"`
	} `json:"timedSchedule"`
}

type Day struct {
	Day     string   `json:"day"`
	Periods []Period `json:"periods"`
}

type Period struct {
	IsCancelled  bool    `json:"isCancelled"`
	PeriodType   string  `json:"periodType"`
	PeriodName   string  `json:"periodName"`
	StartTime    string  `json:"startTime"`
	HeatSetPoint float64 `json:"heatSetPoint"`
	CoolSetPoint float64 `json:"coolSetPoint"`
	FanSwitch    struct {
		Position string `json:"position"`
	} `json:"fanSwitch"`
	Priority struct {
		PriorityType  string `json:"priorityType"`
		SelectedRooms []int  `json:"selectedRooms"`
	} `json:"priority"`
}
