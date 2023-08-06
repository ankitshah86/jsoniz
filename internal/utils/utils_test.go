package utils

import (
	"os"
	"testing"
)

func TestGetEnvBool(t *testing.T) {

	// Set up environment variables
	os.Setenv("ENABLE_FEATURE", "true")
	os.Setenv("DISABLE_FEATURE", "false")

	// Clean up environment variables after the test
	defer func() {
		os.Unsetenv("ENABLE_FEATURE")
		os.Unsetenv("DISABLE_FEATURE")
	}()

	type args struct {
		key          string
		defaultValue bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "key exists, value is true",
			args: args{
				key:          "ENABLE_FEATURE",
				defaultValue: false,
			},
			want: true,
		},
		{
			name: "key exists, value is false",
			args: args{
				key:          "DISABLE_FEATURE",
				defaultValue: true,
			},
			want: false,
		},
		{
			name: "key does not exist, default value is true",
			args: args{
				key:          "NON_EXISTENT_KEY",
				defaultValue: true,
			},
			want: true,
		},
		{
			name: "key does not exist, default value is false",
			args: args{
				key:          "NON_EXISTENT_KEY",
				defaultValue: false,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvBool(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetEnvBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvString(t *testing.T) {

	// Set up environment variables
	os.Setenv("DATABASE_URL", "postgres://user:password@localhost:5432/mydb")
	os.Setenv("API_KEY", " ")

	// Clean up environment variables after the test
	defer func() {
		os.Unsetenv("DATABASE_URL")
		os.Unsetenv("API_KEY")
	}()

	type args struct {
		key          string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "key exists, value is non-empty string",
			args: args{
				key:          "DATABASE_URL",
				defaultValue: "localhost:5432/mydb",
			},
			want: "postgres://user:password@localhost:5432/mydb",
		},
		{
			name: "key exists, value is empty string",
			args: args{
				key:          "API_KEY",
				defaultValue: "default_key",
			},
			want: " ",
		},
		{
			name: "key does not exist, default value is non-empty string",
			args: args{
				key:          "NON_EXISTENT_KEY",
				defaultValue: "default_value",
			},
			want: "default_value",
		},
		{
			name: "key does not exist, default value is empty string",
			args: args{
				key:          "NON_EXISTENT_KEY",
				defaultValue: "",
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvString(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetEnvString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvInt(t *testing.T) {

	// Set up environment variables
	os.Setenv("MAX_CONNECTIONS", "100")
	os.Setenv("MIN_CONNECTIONS", "-100")

	// Clean up environment variables after the test
	defer func() {
		os.Unsetenv("MAX_CONNECTIONS")
		os.Unsetenv("MIN_CONNECTIONS")
	}()

	type args struct {
		key          string
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "key exists, value is positive integer",
			args: args{
				key:          "MAX_CONNECTIONS",
				defaultValue: 10,
			},
			want: 100,
		},
		{
			name: "key exists, value is negative integer",
			args: args{
				key:          "MIN_CONNECTIONS",
				defaultValue: -10,
			},
			want: -100,
		},
		{
			name: "key does not exist, default value is positive integer",
			args: args{
				key:          "NON_EXISTENT_KEY",
				defaultValue: 100,
			},
			want: 100,
		},
		{
			name: "key does not exist, default value is negative integer",
			args: args{
				key:          "NON_EXISTENT_KEY",
				defaultValue: -100,
			},
			want: -100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvInt(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetEnvInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
