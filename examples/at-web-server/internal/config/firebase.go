package config

import "errors"

// Firebase は Firebase の情報を表します。
type Firebase struct {
	// Firebase の証明書情報。
	Credentials string `toml:"credentials"`
	// Web API キー。
	APIKey string `toml:"api_key"`
	// 認証ドメイン。
	AuthDomain string `toml:"auth_domain"`
	// プロジェクトID。
	ProjectID string `toml:"project_id"`
	// ストレージバケット。
	StorageBucket string `toml:"storage_bucket"`
	// プロジェクト番号
	MessagingSenderID string `toml:"messaging_sender_id"`
	// アプリ ID
	AppID string `toml:"app_id"`
}

var defaultFirebase = Firebase{}

func (firebase *Firebase) Validate() error {
	if firebase.Credentials == "" {
		return errors.New("credentials is not present")
	}
	if firebase.APIKey == "" {
		return errors.New("api_key is not present")
	}
	if firebase.AuthDomain == "" {
		return errors.New("auth_domain is not present")
	}
	if firebase.ProjectID == "" {
		return errors.New("project_id is not present")
	}
	if firebase.StorageBucket == "" {
		return errors.New("storage_bucket is not present")
	}
	if firebase.MessagingSenderID == "" {
		return errors.New("messaging_sender_id is not present")
	}
	if firebase.AppID == "" {
		return errors.New("app_id is not present")
	}

	return nil
}

func (firebase *Firebase) Merge(other *Firebase) {
	if firebase == nil || other == nil {
		return
	}

	if other.Credentials != "" {
		firebase.Credentials = other.Credentials
	}
	if other.APIKey != "" {
		firebase.APIKey = other.APIKey
	}
	if other.AuthDomain != "" {
		firebase.AuthDomain = other.AuthDomain
	}
	if other.ProjectID != "" {
		firebase.ProjectID = other.ProjectID
	}
	if other.StorageBucket != "" {
		firebase.StorageBucket = other.StorageBucket
	}
	if other.MessagingSenderID != "" {
		firebase.MessagingSenderID = other.MessagingSenderID
	}
	if other.AppID != "" {
		firebase.AppID = other.AppID
	}
}
