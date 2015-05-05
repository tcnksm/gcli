package skeleton

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		in      *Flag
		success bool
	}{
		{
			in: &Flag{
				LongName:    "debug",
				TypeString:  "bool",
				Description: "Run as a DEBUG mode",
			},

			success: true,
		},

		{
			in: &Flag{
				LongName:   "debug",
				TypeString: "bool",
			},
			success: true,
		},

		{
			in: &Flag{
				LongName: "debug",
			},
			success: false,
		},

		{
			in:      &Flag{},
			success: false,
		},
	}

	for i, tt := range tests {

		err := tt.in.Validate()
		if err != nil && !tt.success {
			continue
		}

		if err == nil && !tt.success {
			t.Fatalf("#%d expect Validate to fail", i)
		}

		if err != nil {
			t.Fatalf("#%d expect Fix not to fail but %q", i, err.Error())
		}
	}
}

func TestFix(t *testing.T) {
	tests := []struct {
		in      *Flag
		exp     *Flag
		success bool
	}{
		{
			in: &Flag{
				LongName:    "debug",
				TypeString:  "bool",
				Description: "Run as DEBUG mode",
			},
			exp: &Flag{
				Name:        "Debug",
				ShortName:   "d",
				LongName:    "debug",
				TypeString:  "Bool",
				Default:     false,
				Description: "Run as DEBUG mode",
			},
			success: true,
		},

		{
			in: &Flag{
				LongName:    "token",
				TypeString:  "s",
				Description: "",
			},
			exp: &Flag{
				Name:        "Token",
				ShortName:   "t",
				LongName:    "token",
				TypeString:  "String",
				Default:     "\"\"",
				Description: "",
			},
			success: true,
		},
	}

	for i, tt := range tests {

		err := tt.in.Fix()
		if err != nil && !tt.success {
			continue
		}

		if err == nil && !tt.success {
			t.Fatalf("#%d expect Fix to fail", i)
		}

		if err != nil {
			t.Fatalf("#%d expect Fix not to fail but %q", i, err.Error())
		}

		if *tt.in != *tt.exp {
			t.Errorf("#%d expect %v to eq %v", i, tt.in, tt.exp)
		}
	}
}

func TestFixTypeString(t *testing.T) {

	tests := []struct {
		in            *Flag
		success       bool
		expTypeString string
		expDefault    interface{}
	}{
		{
			in:            &Flag{TypeString: "int"},
			success:       true,
			expTypeString: "Int",
			expDefault:    0,
		},

		{
			in:            &Flag{TypeString: "Int"},
			success:       true,
			expTypeString: "Int",
			expDefault:    0,
		},

		{
			in:            &Flag{TypeString: "i"},
			success:       true,
			expTypeString: "Int",
			expDefault:    0,
		},

		{
			in:            &Flag{TypeString: "string"},
			success:       true,
			expTypeString: "String",
			expDefault:    "\"\"",
		},

		{
			in:            &Flag{TypeString: "s"},
			success:       true,
			expTypeString: "String",
			expDefault:    "\"\"",
		},

		{
			in:            &Flag{TypeString: "str"},
			success:       true,
			expTypeString: "String",
			expDefault:    "\"\"",
		},

		{
			in:            &Flag{TypeString: "bool"},
			success:       true,
			expTypeString: "Bool",
			expDefault:    false,
		},

		{
			in:            &Flag{TypeString: "b"},
			success:       true,
			expTypeString: "Bool",
			expDefault:    false,
		},

		{
			in:            &Flag{TypeString: "enexpected_type"},
			success:       false,
			expTypeString: "Bool",
			expDefault:    false,
		},
	}

	for i, tt := range tests {

		err := tt.in.fixTypeString()
		if err != nil && !tt.success {
			continue
		}

		if err == nil && !tt.success {
			t.Fatalf("#%d expect fixTypeString to fail", i)
		}

		if err != nil {
			t.Fatalf("#%d expect fixTypeString not to fail but %q", i, err.Error())
		}

		if tt.in.TypeString != tt.expTypeString {
			t.Errorf("#%d expect %q to eq %q", i, tt.in.TypeString, tt.expTypeString)
		}

		if tt.in.Default != tt.expDefault {
			t.Errorf("#%d expect %v to eq %v", i, tt.in.Default, tt.expDefault)
		}
	}
}
