package upcloud

import (
	"context"
	"fmt"
	"net/http"
)

// AccountService handles communication with account related methods of the Upcloud API
// https://developers.upcloud.com/1.3/3-accounts/
type AccountService service

// AccountInformation represents the current account limits
// Either ResourceLimits or TrailResourceLimits depending on account, check for nil
type AccountInformation struct {
	Credits  float64 `json:"credits"`
	Username string  `json:"username"`

	ResourceLimits *struct {
		Cores               int `json:"cores"`                 //Maximum number of CPU cores
		DetachedFloatingIPs int `json:"detached_floating_ips"` //Maximum number of detached floating IP addresses
		Memory              int `json:"memory"`                //Maximum amount of memory in MiB
		Networks            int `json:"networks"`              //Maximum number of networks
		PublicIPv4          int `json:"public_ipv4"`           //Maximum number of networks
		PublicIPv6          int `json:"public_ipv6"`           //Maximum number of IPv6 addresses
		StorageHDD          int `json:"storage_hdd"`           //Maximum amount of HDD storage space in MiB
		StorageSDD          int `json:"storage_sdd"`           //Maximum amount of SSD storage space in MiB
	} `json:"resource_limits,omitempty"`

	TrialResourceLimits *struct {
		FirewallRestrictions     int    `json:"trial_firewall_restrictions"`       //If 1, firewall option is disabled
		PeriodLength             int    `json:"trial_period_length"`               //Trial period length in hours
		ServerMaxCores           int    `json:"trial_server_max_cores"`            //Maximum number of CPU cores per server
		ServerMaxMemory          int    `json:"trial_server_max_memory"`           //Maximum amount of memory in MiB per server
		ServerMaxPublicIPv4      int    `json:"trial_server_max_public_ipv4"`      //Maximum number of public IPv4 addresses per server
		ServerMaxPublicIPv6      int    `json:"trial_server_max_public_ipv6"`      //Maximum number of public IPv6 addresses per server
		StorageMaxSize           int    `json:"trial_storage_max_size"`            //Maximum storage size in GiB
		StorageTier              string `json:"trial_storage_tier"`                //Storage tier type
		TotalDetachedFloatingIPs int    `json:"trial_total_detached_floating_ips"` //Maximum number of detached floating IP addresses
		TotalNetworks            int    `json:"trial_total_networks"`              //Maximum number of networks
		TotalPublicIPv4          int    `json:"trial_total_public_ipv4"`           //Maximum number of public IPv4 addresses
		TotalPublicIPv6          int    `json:"trial_total_public_ipv6"`           //Maximum number of public IPv6 addresses
		TotalServerCores         int    `json:"trial_total_server_cores"`          //Maximum number of CPU cores
		TotalServerMemory        int    `json:"trial_total_server_memory"`         //Maximum amount of memory in GiB
		TotalServers             int    `json:"trial_total_servers"`               //Maximum number of servers
		TotalStorageSize         int    `json:"trial_total_storage_size"`          //Maximum amount of storage in GiB
		TotalStorages            int    `json:"trial_total_storages"`              //Maximum number of storage devices
		UserDetachedFloatingIPs  int    `json:"user_detached_floating_ips"`        //Number of detached floating IP addresses
		UserNetworks             int    `json:"user_networks"`                     //Number of networks in use
		UserPublicIPv4           int    `json:"user_public_ipv4"`                  //Number of public IPv4 addresses in use
		UserPublicIPv6           int    `json:"user_public_ipv6"`                  //Number of public IPv6 addresses in use
		UserServerCores          int    `json:"user_server_cores"`                 //Number of CPU cores in use
		UserServerMemory         int    `json:"user_server_memory"`                //Amount of memory in use MiB
	} `json:"trial_resource_limits,omitempty"`
}

// AccountInformationResponse represents the response from account information API methods
type AccountInformationResponse struct {
	AccountInformation *AccountInformation `json:"account"`
}

// GetAccountInformation returns the credits and limits on the account.
// https://developers.upcloud.com/1.3/3-accounts/#get-account-information
func (s *AccountService) GetAccountInformation(ctx context.Context) (*AccountInformation, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "account", nil)
	if err != nil {
		return nil, nil, err
	}

	accountInformation := new(AccountInformation)
	resp, err := s.client.Do(ctx, req, &AccountInformationResponse{AccountInformation: accountInformation})
	if err != nil {
		return nil, resp, err
	}

	return accountInformation, resp, nil
}

// Roles represents the list of roles an account may have
type Roles struct {
	Role []string `json:"role"`
}

// AccountListEntry represents an entry in the account list
type AccountListEntry struct {
	Roles    *Roles `json:"roles"`
	Type     string `json:"type"`
	Username string `json:"username"`
}

// AccountList represents the list of accounts
type AccountList struct {
	Accounts []AccountListEntry `json:"account"`
}

// AccountListResponse represents the response from the list account API methods.
type AccountListResponse struct {
	Accounts *AccountList `json:"accounts"`
}

// ListAccounts returns the list of accounts
// https://developers.upcloud.com/1.3/3-accounts/#get-account-list
func (s *AccountService) ListAccounts(ctx context.Context) (*AccountList, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "account/list", nil)
	if err != nil {
		return nil, nil, err
	}

	accountList := new(AccountList)
	resp, err := s.client.Do(ctx, req, &AccountListResponse{Accounts: accountList})
	if err != nil {
		return nil, resp, err
	}

	return accountList, resp, nil
}

// Account represents a detailed account on Upcloud
type Account struct {
	MainAccount string `json:"main_account,omitempty"`
	Type        string `json:"type,omitempty"`
	Username    string `json:"username,omitempty"` // Implied by URL on modify
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Company     string `json:"company,omitempty"`
	Address     string `json:"address,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	Country     string `json:"country,omitempty"`
	Currency    string `json:"currency"` //EUR/GBP/USD/SGD
	Language    string `json:"language"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	VATNumber   string `json:"vat_number,omitempty"`
	Timezone    string `json:"timezone"`
	Password    string `json:"password,omitempty"`
	// Campaigns?
	Roles                  Roles  `json:"roles,omitempty"`
	AllowAPI               string `json:"allow_api"`                           //yes/no
	AllowGUI               string `json:"allow_gui"`                           //yes/no
	Enable3rdPartyServices string `json:"enable_3rd_party_services,omitempty"` //yes/no

	NetworkAccess struct {
		Networks []string `json:"network,omitempty"`
	} `json:"network_access"`

	ServerAccess struct {
		Servers []struct {
			Storage string `json:"storage"`
			UUID    string `json:"uuid"`
		} `json:"server"`
	} `json:"server_access"`

	StorageAccess struct {
		Storage []string `json:"storage"`
	} `json:"storage_access"`

	TagAccess struct {
		Tags []struct {
			Name    string `json:"name"`
			Storage string `json:"storage"` //yes/no
		} `json:"tag"`
	} `json:"tag_access"`

	IPFilters struct {
		IPFilters []string `json:"ip_filter"`
	} `json:"ip_filters,omitempty"`
}

// AccountDetails represents the request/response from the Details API method
type AccountDetails struct {
	Account *Account `json:"account"`
}

// GetAccountDetails returns a detailed account response for a given username
// https://developers.upcloud.com/1.3/3-accounts/#get-account-details
func (s *AccountService) GetAccountDetails(ctx context.Context, username string) (*Account, *http.Response, error) {
	u := fmt.Sprintf("account/details/%v", username)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	account := new(Account)
	resp, err := s.client.Do(ctx, req, &AccountDetails{Account: account})
	if err != nil {
		return nil, resp, err
	}

	return account, resp, nil
}

// ModifyAccount you probably want ModifyAccountDetails or ModifySubAccountDetails
// kind can only be `details` or `sub`.
// https://developers.upcloud.com/1.3/3-accounts/#modify-account-details
func (s *AccountService) ModifyAccount(ctx context.Context, kind string, account *Account, username string) (*http.Response, error) {
	trimAccount := *account
	trimAccount.MainAccount = ""
	trimAccount.Username = ""
	trimAccount.Type = ""
	trimAccount.Password = ""
	u := fmt.Sprintf("account/%v/%v", kind, username)
	req, err := s.client.NewRequest("PUT", u, &AccountDetails{Account: &trimAccount})
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// ModifyAccountDetails modifies the details of the given username with properties defined in acc.
func (s *AccountService) ModifyAccountDetails(ctx context.Context, acc *Account, username string) (*http.Response, error) {
	return s.ModifyAccount(ctx, "details", acc, username)
}

// ModifySubAccountDetails modifies the details of a sub account of the given username with properties defined in acc.
func (s *AccountService) ModifySubAccountDetails(ctx context.Context, acc *Account, username string) (*http.Response, error) {
	// MainAccount, Type and Username must not be set
	return s.ModifyAccount(ctx, "sub", acc, username)
}

// SubAccount represents the request/response wrapper for AddSubAccount
type SubAccount struct {
	Account *Account `json:"sub_account"`
}

// AddSubAccount creates a new sub account with the details provided in acc.
// https://developers.upcloud.com/1.3/3-accounts/#add-subaccount
func (s *AccountService) AddSubAccount(ctx context.Context, acc *Account) (*http.Response, error) {
	req, err := s.client.NewRequest("POST", "account/sub", &SubAccount{Account: acc})
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeleteSubAccount deletes a sub account of the given username
// https://developers.upcloud.com/1.3/3-accounts/#delete-subaccount
func (s *AccountService) DeleteSubAccount(ctx context.Context, username string) (*http.Response, error) {
	u := fmt.Sprintf("account/sub/%v", username)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
