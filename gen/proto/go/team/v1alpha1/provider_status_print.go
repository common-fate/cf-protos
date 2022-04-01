package teamv1alpha1

func (p ProviderStatus) DisplayString() string {
	switch p {
	case ProviderStatus_PROVIDER_STATUS_PENDING:
		return "PENDING"
	case ProviderStatus_PROVIDER_STATUS_READY:
		return "READY"
	}
	return "UNSPECIFIED"
}
