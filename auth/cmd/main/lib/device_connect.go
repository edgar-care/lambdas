package lib

import (
	"encoding/json"
	"fmt"
	edgarlib "github.com/edgar-care/edgarlib/double_auth"
	"net"
	"net/http"
	"strings"
	"time"
)

type IPInfoResponse struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

type DeviceInfo struct {
	IPAddress     string  `json:"ip_address"`
	DeviceName    string  `json:"device_name"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	OperationTime string  `json:"operation_time"`
}

func GetIPAddress(r *http.Request) string {
	// Vérifie l'en-tête X-Forwarded-For
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		// X-Forwarded-For peut contenir une liste d'adresses IP séparées par des virgules
		ips := strings.Split(ip, ",")
		for _, addr := range ips {
			addr = strings.TrimSpace(addr)
			// Valide si l'IP est une adresse IPv4 valide
			if net.ParseIP(addr) != nil && strings.Contains(addr, ".") {
				return addr
			}
		}
	}

	// Vérifie l'en-tête X-Real-IP
	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		// Valide si l'IP est une adresse IPv4 valide
		if net.ParseIP(ip) != nil && strings.Contains(ip, ".") {
			return ip
		}
	}

	// Utilise l'adresse IP distante si aucune des autres méthodes n'a fonctionné
	ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	if net.ParseIP(ip) != nil && strings.Contains(ip, ".") {
		return ip
	}

	return ""
}

func getGeoLocationAndDeviceName(ip string) (string, float64, float64, error) {
	url := fmt.Sprintf("https://ipinfo.io/%s/json?token=bde9bcbd8cf979", ip)

	resp, err := http.Get(url)
	if err != nil {
		return "Unknown Device", 0.0, 0.0, err
	}
	defer resp.Body.Close()

	var ipInfo IPInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&ipInfo); err != nil {
		return "Unknown Device", 0.0, 0.0, err
	}

	var latitude, longitude float64
	if _, err := fmt.Sscanf(ipInfo.Loc, "%f,%f", &latitude, &longitude); err != nil {
		return "Unknown Device", 0.0, 0.0, err
	}

	deviceName := ipInfo.Hostname
	if deviceName == "" {
		deviceName = "Unknown Device"
	}

	return deviceName, latitude, longitude, nil
}

func GetDeviceInfo(w http.ResponseWriter, r *http.Request, patientId string) string {

	owner_id := AuthMiddleware(patientId) //"667928fd912f76ecb2fd403e"

	ip := "62.34.233.53" // GetIPAddress(r)
	deviceName, latitude, longitude, err := getGeoLocationAndDeviceName(ip)
	if err != nil {
		fmt.Println("Error fetching double_auth information:", err)
		http.Error(w, "Error fetching double_auth information", 400)
		return ""
	}

	operationTime := time.Now().Unix() // Unix timestamp

	input := edgarlib.CreateDeviceConnectInput{
		DeviceName: deviceName,
		Ip:         ip,
		Latitude:   latitude,
		Longitude:  longitude,
		Date:       int(operationTime),
	}

	response := edgarlib.CreateDeviceConnect(input, owner_id)
	if response.Err != nil {
		WriteError(w, response.Code, response.Err.Error())
		return ""
	}
	return "Successfully fetched double_auth information"
}
