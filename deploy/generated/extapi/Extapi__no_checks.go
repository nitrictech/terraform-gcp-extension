//go:build no_runtime_type_checking

package extapi

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Extapi) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (e *jsiiProxy_Extapi) validateAddProviderParameters(provider interface{}) error {
	return nil
}

func (e *jsiiProxy_Extapi) validateGetStringParameters(output *string) error {
	return nil
}

func (e *jsiiProxy_Extapi) validateInterpolationForOutputParameters(moduleOutput *string) error {
	return nil
}

func (e *jsiiProxy_Extapi) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateExtapi_IsConstructParameters(x interface{}) error {
	return nil
}

func validateExtapi_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Extapi) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Extapi) validateSetOpenapiSpecParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Extapi) validateSetStackIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Extapi) validateSetTargetServicesParameters(val *map[string]*string) error {
	return nil
}

func validateNewExtapiParameters(scope constructs.Construct, id *string, config *ExtapiConfig) error {
	return nil
}

