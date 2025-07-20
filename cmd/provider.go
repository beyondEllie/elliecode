package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

const (
	ProviderConfigDirName  = "ellie"
	ProviderConfigFileName = ".ellie.env"
)

var providerSetCmd = &cobra.Command{
	Use:   "provider set",
	Short: "Set provider and API key",
	RunE: func(cmd *cobra.Command, args []string) error {
		provider, _ := cmd.Flags().GetString("provider")
		apiKey, _ := cmd.Flags().GetString("api-key")
		if provider == "" || apiKey == "" {
			return fmt.Errorf("Both --provider and --api-key are required")
		}

		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("Unable to determine user home directory: %w", err)
		}
		configDir := filepath.Join(homeDir, ProviderConfigDirName)
		configPath := filepath.Join(configDir, ProviderConfigFileName)

		if err := os.MkdirAll(configDir, 0700); err != nil {
			return fmt.Errorf("Failed to create config directory: %w", err)
		}

		// Load existing config if present
		config := map[string]string{}
		if _, err := os.Stat(configPath); err == nil {
			config, _ = godotenv.Read(configPath)
		}

		// Set provider and API key
		config[fmt.Sprintf("%s_API_KEY", providerToEnv(provider))] = apiKey

		if err := godotenv.Write(config, configPath); err != nil {
			return fmt.Errorf("Failed to write config file: %w", err)
		}
		if err := os.Chmod(configPath, 0600); err != nil {
			fmt.Printf("Warning: Failed to set secure permissions on config file: %v\n", err)
		}
		fmt.Printf("✅ Provider '%s' API key set in %s\n", provider, configPath)
		return nil
	},
}

func providerToEnv(provider string) string {
	switch provider {
	case "openai":
		return "OPENAI"
	case "anthropic":
		return "ANTHROPIC"
	case "gemini":
		return "GEMINI"
	case "groq":
		return "GROQ"
	case "openrouter":
		return "OPENROUTER"
	case "copilot":
		return "GITHUB"
	case "azure":
		return "AZURE_OPENAI"
	case "bedrock":
		return "AWS"
	case "vertexai":
		return "VERTEXAI"
	case "xai":
		return "XAI"
	default:
		return provider
	}
}

func init() {
	providerSetCmd.Flags().String("provider", "", "Provider name (e.g. openai, anthropic, gemini, etc.)")
	providerSetCmd.Flags().String("api-key", "", "API key for the provider")
	providerSetCmd.MarkFlagRequired("provider")
	providerSetCmd.MarkFlagRequired("api-key")
	rootCmd.AddCommand(providerSetCmd)
}
