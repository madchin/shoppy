package ports

import "backend/internal/users/domain/user"

func mapDomainAddressesToResponseAddresses(domainAddresses user.Addresses) Addresses {
	var addresses Addresses
	for _, address := range domainAddresses {
		address := Address{City: address.City(), Country: address.Country(), PostalCode: address.PostalCode(), Street: address.Street()}
		addresses = append(addresses, address)
	}
	return addresses
}
