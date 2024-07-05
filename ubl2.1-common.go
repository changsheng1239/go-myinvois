package main

// A character string (i.e. a finite set of characters), generally in the form of words of a
// language.
//
// A character string (i.e. a finite set of characters) generally in the form of words of a
// language.
//
// A character string that constitutes the distinctive designation of a person, place, thing
// or concept.
type TextType struct {
	Empty            string  `json:"_"`
	LanguageID       *string `json:"languageID,omitempty"`
	LanguageLocaleID *string `json:"languageLocaleID,omitempty"`
}

// A character string (letters, figures, or symbols) that for brevity and/or language
// independence may be used to represent or replace a definitive value or text of an
// attribute, together with relevant supplementary information.
//
// A character string (letters, figures, or symbols) that for brevity and/or languange
// independence may be used to represent or replace a definitive value or text of an
// attribute together with relevant supplementary information.
type CodeType struct {
	Empty          string  `json:"_"`
	LanguageID     *string `json:"languageID,omitempty"`
	ListAgencyID   *string `json:"listAgencyID,omitempty"`
	ListAgencyName *string `json:"listAgencyName,omitempty"`
	ListID         *string `json:"listID,omitempty"`
	ListName       *string `json:"listName,omitempty"`
	ListSchemeURI  *string `json:"listSchemeURI,omitempty"`
	ListURI        *string `json:"listURI,omitempty"`
	ListVersionID  *string `json:"listVersionID,omitempty"`
	Name           *string `json:"name,omitempty"`
}

// A class to describe a customer party.
type CustomerPartyDetails struct {
	// A customer contact for accounting.
	AccountingContact []ContactDetails `json:"AccountingContact,omitempty"`
	// An identifier for the customer's account, assigned by a third party.
	AdditionalAccountID []IdentifierType `json:"AdditionalAccountID,omitempty"`
	// A customer contact for purchasing.
	BuyerContact []ContactDetails `json:"BuyerContact,omitempty"`
	// An identifier for the customer's account, assigned by the customer itself.
	CustomerAssignedAccountID []IdentifierType `json:"CustomerAssignedAccountID,omitempty"`
	// A customer contact for deliveries.
	DeliveryContact []ContactDetails `json:"DeliveryContact,omitempty"`
	// The customer party itself.
	Party []PartyDetails `json:"Party,omitempty"`
	// An identifier for the customer's account, assigned by the supplier.
	SupplierAssignedAccountID []IdentifierType `json:"SupplierAssignedAccountID,omitempty"`
}

// A class to describe a contactable person or department in an organization.
type ContactDetails struct {
	// The primary email address of this contact.
	ElectronicMail []TextType `json:"ElectronicMail,omitempty"`
	// An identifier for this contact.
	ID []IdentifierType `json:"ID,omitempty"`
	// The name of this contact. It is recommended that this be used for a functional name and
	// not a personal name.
	Name []TextType `json:"Name,omitempty"`
	// Free-form text conveying information that is not contained explicitly in other
	// structures; in particular, a textual description of the circumstances under which this
	// contact can be used (e.g., "emergency" or "after hours").
	Note []TextType `json:"Note,omitempty"`
	// Another means of communication with this contact.
	OtherCommunication []CommunicationDetails `json:"OtherCommunication,omitempty"`
	// The primary fax number of this contact.
	Telefax []TextType `json:"Telefax,omitempty"`
	// The primary telephone number of this contact.
	Telephone []TextType `json:"Telephone,omitempty"`
}

// A character string to identify and uniquely distinguish one instance of an object in an
// identification scheme from all other objects in the same scheme, together with relevant
// supplementary information.
//
// A character string to identify and distinguish uniquely, one instance of an object in an
// identification scheme from all other objects in the same scheme together with relevant
// supplementary information.
type IdentifierType struct {
	Empty            string  `json:"_"`
	SchemeAgencyID   *string `json:"schemeAgencyID,omitempty"`
	SchemeAgencyName *string `json:"schemeAgencyName,omitempty"`
	SchemeDataURI    *string `json:"schemeDataURI,omitempty"`
	SchemeID         *string `json:"schemeID,omitempty"`
	SchemeName       *string `json:"schemeName,omitempty"`
	SchemeURI        *string `json:"schemeURI,omitempty"`
	SchemeVersionID  *string `json:"schemeVersionID,omitempty"`
}

// A class to describe a means of communication.
type CommunicationDetails struct {
	// The method of communication, expressed as text.
	Channel []TextType `json:"Channel,omitempty"`
	// The method of communication, expressed as a code.
	ChannelCode []CodeType `json:"ChannelCode,omitempty"`
	// An identifying value (phone number, email address, etc.) for this channel of communication
	Value []TextType `json:"Value,omitempty"`
}

// A class to describe a party contracting to provide services, such as transportation,
// finance, etc.
type ServiceProviderPartyDetails struct {
	// An identifier for this service provider.
	ID []IdentifierType `json:"ID,omitempty"`
	// The party providing the service.
	Party []PartyDetails `json:"Party"`
	// The contact for the service provider.
	SellerContact []ContactDetails `json:"SellerContact,omitempty"`
	// The type of service provided, expressed as text.
	ServiceType []TextType `json:"ServiceType,omitempty"`
	// The type of service provided, expressed as a code.
	ServiceTypeCode []CodeType `json:"ServiceTypeCode,omitempty"`
}

// A class to describe a power of attorney.
type PowerOfAttorneyDetails struct {
	// The party who acts as an agent or fiduciary for the principal and who holds this power of
	// attorney on behalf of the principal.
	AgentParty []PartyDetails `json:"AgentParty"`
	// Text describing this power of attorney.
	Description []TextType `json:"Description,omitempty"`
	// An identifier for this power of attorney.
	ID []IdentifierType `json:"ID,omitempty"`
	// The date on which this power of attorney was issued.
	IssueDate []DateType `json:"IssueDate,omitempty"`
	// The time at which this power of attorney was issued.
	IssueTime []TimeType `json:"IssueTime,omitempty"`
	// A reference to a mandate associated with this power of attorney.
	MandateDocumentReference []DocumentReferenceDetails `json:"MandateDocumentReference,omitempty"`
	// The party notarizing this power of attorney.
	NotaryParty []PartyDetails `json:"NotaryParty,omitempty"`
	// An association to a WitnessParty.
	WitnessParty []PartyDetails `json:"WitnessParty,omitempty"`
}

// A class to describe the result of an attempt to verify a signature.
type ResultOfVerificationDetails struct {
	// The signing party.
	SignatoryParty []PartyDetails `json:"SignatoryParty,omitempty"`
	// The verification process.
	ValidateProcess []TextType `json:"ValidateProcess,omitempty"`
	// The tool used to verify the signature.
	ValidateTool []TextType `json:"ValidateTool,omitempty"`
	// The version of the tool used to verify the signature.
	ValidateToolVersion []TextType `json:"ValidateToolVersion,omitempty"`
	// The date upon which verification took place.
	ValidationDate []DateType `json:"ValidationDate,omitempty"`
	// A code signifying the result of the verification.
	ValidationResultCode []CodeType `json:"ValidationResultCode,omitempty"`
	// The time at which verification took place.
	ValidationTime []TimeType `json:"ValidationTime,omitempty"`
	// An identifier for the organization, person, service, or server that verified the
	// signature.
	ValidatorID []IdentifierType `json:"ValidatorID,omitempty"`
}

// A class to define a reference to a document.
type DocumentReferenceDetails struct {
	// The referenced document as an attachment to the document from which it is referenced.
	Attachment []AttachmentDetails `json:"Attachment,omitempty"`
	// An indicator that the referenced document is a copy (true) or the original (false).
	CopyIndicator []IndicatorType `json:"CopyIndicator,omitempty"`
	// Text describing the referenced document.
	DocumentDescription []TextType `json:"DocumentDescription,omitempty"`
	// A code signifying the status of the reference document with respect to its original state.
	DocumentStatusCode []CodeType `json:"DocumentStatusCode,omitempty"`
	// The type of document being referenced, expressed as text.
	DocumentType []TextType `json:"DocumentType,omitempty"`
	// The type of document being referenced, expressed as a code.
	DocumentTypeCode []CodeType `json:"DocumentTypeCode,omitempty"`
	// An identifier for the referenced document.
	ID []IdentifierType `json:"ID"`
	// The date, assigned by the sender of the referenced document, on which the document was
	// issued.
	IssueDate []DateType `json:"IssueDate,omitempty"`
	// The party who issued the referenced document.
	IssuerParty []PartyDetails `json:"IssuerParty,omitempty"`
	// The time, assigned by the sender of the referenced document, at which the document was
	// issued.
	IssueTime []TimeType `json:"IssueTime,omitempty"`
	// An identifier for the language used in the referenced document.
	LanguageID []IdentifierType `json:"LanguageID,omitempty"`
	// A code signifying the locale in which the language in the referenced document is used.
	LocaleCode []CodeType `json:"LocaleCode,omitempty"`
	// The result of an attempt to verify a signature associated with the referenced document.
	ResultOfVerification []ResultOfVerificationDetails `json:"ResultOfVerification,omitempty"`
	// A universally unique identifier for this document reference.
	UUID []IdentifierType `json:"UUID,omitempty"`
	// The period for which this document reference is valid.
	ValidityPeriod []PeriodDetails `json:"ValidityPeriod,omitempty"`
	// An identifier for the current version of the referenced document.
	VersionID []IdentifierType `json:"VersionID,omitempty"`
	// A reference to another place in the same XML document instance in which DocumentReference
	// appears.
	XPath []TextType `json:"XPath,omitempty"`
}

// A class to describe a person.
type PersonDetails struct {
	// This person's date of birth.
	BirthDate []DateType `json:"BirthDate,omitempty"`
	// The name of the place where this person was born, expressed as text.
	BirthplaceName []TextType `json:"BirthplaceName,omitempty"`
	// Contact information for this person.
	Contact []ContactDetails `json:"Contact,omitempty"`
	// This person's family name.
	FamilyName []TextType `json:"FamilyName,omitempty"`
	// The financial account associated with this person.
	FinancialAccount []FinancialAccountDetails `json:"FinancialAccount,omitempty"`
	// This person's given name.
	FirstName []TextType `json:"FirstName,omitempty"`
	// A code (e.g., ISO 5218) signifying the gender of this person.
	GenderCode []CodeType `json:"GenderCode,omitempty"`
	// An identifier for this person.
	ID []IdentifierType `json:"ID,omitempty"`
	// A reference to a document that can precisely identify this person (e.g., a driver's
	// license).
	IdentityDocumentReference []DocumentReferenceDetails `json:"IdentityDocumentReference,omitempty"`
	// This person's job title (for a particular role) within an organization.
	JobTitle []TextType `json:"JobTitle,omitempty"`
	// This person's middle name(s) or initials.
	MiddleName []TextType `json:"MiddleName,omitempty"`
	// A suffix to this person's name (e.g., PhD, OBE, Jr).
	NameSuffix []TextType `json:"NameSuffix,omitempty"`
	// An identifier for this person's nationality.
	NationalityID []IdentifierType `json:"NationalityID,omitempty"`
	// The department or subdivision of an organization that this person belongs to (in a
	// particular role).
	OrganizationDepartment []TextType `json:"OrganizationDepartment,omitempty"`
	// This person's second family name.
	OtherName []TextType `json:"OtherName,omitempty"`
	// This person's address of residence.
	ResidenceAddress []AddressDetails `json:"ResidenceAddress,omitempty"`
	// This person's title of address (e.g., Mr, Ms, Dr, Sir).
	Title []TextType `json:"Title,omitempty"`
}

// A class to describe a shareholder party.
type ShareholderPartyDetails struct {
	// The shareholder participation, expressed as a percentage.
	PartecipationPercent []NumericType `json:"PartecipationPercent,omitempty"`
	// The shareholder party.
	Party []PartyDetails `json:"Party,omitempty"`
}

// A class to describe a party as a legal entity.
type PartyLegalEntityDetails struct {
	// An identifier for the party as registered within a company registration scheme.
	CompanyID []IdentifierType `json:"CompanyID,omitempty"`
	// The company legal status, expressed as a text.
	CompanyLegalForm []TextType `json:"CompanyLegalForm,omitempty"`
	// A code signifying the party's legal status.
	CompanyLegalFormCode []CodeType `json:"CompanyLegalFormCode,omitempty"`
	// A code signifying the party's liquidation status.
	CompanyLiquidationStatusCode []CodeType `json:"CompanyLiquidationStatusCode,omitempty"`
	// The corporate registration scheme used to register the party.
	CorporateRegistrationScheme []CorporateRegistrationSchemeDetails `json:"CorporateRegistrationScheme,omitempty"`
	// The number of shares in the capital stock of a corporation.
	CorporateStockAmount []AmountType `json:"CorporateStockAmount,omitempty"`
	// An indicator that all shares of corporate stock have been paid by shareholders (true) or
	// not (false).
	FullyPaidSharesIndicator []IndicatorType `json:"FullyPaidSharesIndicator,omitempty"`
	// The head office of the legal entity
	HeadOfficeParty []PartyDetails `json:"HeadOfficeParty,omitempty"`
	// The registered address of the party within a corporate registration scheme.
	RegistrationAddress []AddressDetails `json:"RegistrationAddress,omitempty"`
	// The registration date of the CompanyID.
	RegistrationDate []DateType `json:"RegistrationDate,omitempty"`
	// The date upon which a registration expires (e.g., registration for an import/export
	// license).
	RegistrationExpirationDate []DateType `json:"RegistrationExpirationDate,omitempty"`
	// The name of the party as registered with the relevant legal authority.
	RegistrationName []TextType `json:"RegistrationName,omitempty"`
	// A party owning shares in this legal entity.
	ShareholderParty []ShareholderPartyDetails `json:"ShareholderParty,omitempty"`
	// An indicator that the company is owned and controlled by one person (true) or not (false).
	SoleProprietorshipIndicator []IndicatorType `json:"SoleProprietorshipIndicator,omitempty"`
}

// A class to describe an organization, sub-organization, or individual fulfilling a role in
// a business process.
type PartyDetails struct {
	// A party who acts as an agent for this party.
	AgentParty []PartyDetails `json:"AgentParty,omitempty"`
	// The primary contact for this party.
	Contact []ContactDetails `json:"Contact,omitempty"`
	// An identifier for the end point of the routing service (e.g., EAN Location Number, GLN).
	EndpointID []IdentifierType `json:"EndpointID,omitempty"`
	// The financial account associated with this party.
	FinancialAccount []FinancialAccountDetails `json:"FinancialAccount,omitempty"`
	// This party's Industry Classification Code.
	IndustryClassificationCode []CodeType `json:"IndustryClassificationCode,omitempty"`
	// The language associated with this party.
	Language []LanguageDetails `json:"Language,omitempty"`
	// An identifier for this party's logo.
	LogoReferenceID []IdentifierType `json:"LogoReferenceID,omitempty"`
	// An indicator that this party is "for the attention of" (FAO) (true) or not (false).
	MarkAttentionIndicator []IndicatorType `json:"MarkAttentionIndicator,omitempty"`
	// An indicator that this party is "care of" (c/o) (true) or not (false).
	MarkCareIndicator []IndicatorType `json:"MarkCareIndicator,omitempty"`
	// An identifier for this party.
	PartyIdentification []PartyIdentificationDetails `json:"PartyIdentification,omitempty"`
	// A description of this party as a legal entity.
	PartyLegalEntity []PartyLegalEntityDetails `json:"PartyLegalEntity,omitempty"`
	// A name for this party.
	PartyName []PartyNameDetails `json:"PartyName,omitempty"`
	// A tax scheme applying to this party.
	PartyTaxScheme []PartyTaxSchemeDetails `json:"PartyTaxScheme,omitempty"`
	// A person associated with this party.
	Person []PersonDetails `json:"Person,omitempty"`
	// The physical location of this party.
	PhysicalLocation []LocationDetails `json:"PhysicalLocation,omitempty"`
	// The party's postal address.
	PostalAddress []AddressDetails `json:"PostalAddress,omitempty"`
	// A power of attorney associated with this party.
	PowerOfAttorney []PowerOfAttorneyDetails `json:"PowerOfAttorney,omitempty"`
	// A party providing a service to this party.
	ServiceProviderParty []ServiceProviderPartyDetails `json:"ServiceProviderParty,omitempty"`
	// The Uniform Resource Identifier (URI) that identifies this party's web site; i.e., the
	// web site's Uniform Resource Locator (URL).
	WebsiteURI []IdentifierType `json:"WebsiteURI,omitempty"`
}

// One calendar day according the Gregorian calendar.
type DateType struct {
	Empty string `json:"_"`
}

// An instance of time that occurs every day.
type TimeType struct {
	Empty string `json:"_"`
}

// A class to describe an attached document. An attachment can refer to an external document
// or be included with the document being exchanged.
type AttachmentDetails struct {
	// A binary large object containing an attached document.
	EmbeddedDocumentBinaryObject []BinaryObjectType `json:"EmbeddedDocumentBinaryObject,omitempty"`
	// A reference to an attached document that is external to the document(s) being exchanged.
	ExternalReference []ExternalReferenceDetails `json:"ExternalReference,omitempty"`
}

// A set of finite-length sequences of binary octets.
type BinaryObjectType struct {
	Empty            string  `json:"_"`
	CharacterSetCode *string `json:"characterSetCode,omitempty"`
	EncodingCode     *string `json:"encodingCode,omitempty"`
	Filename         *string `json:"filename,omitempty"`
	Format           *string `json:"format,omitempty"`
	MIMECode         string  `json:"mimeCode"`
	URI              *string `json:"uri,omitempty"`
}

// A class to describe an external object, such as a document stored at a remote location.
type ExternalReferenceDetails struct {
	// A code signifying the character set of an external document.
	CharacterSetCode []CodeType `json:"CharacterSetCode,omitempty"`
	// Text describing the external object.
	Description []TextType `json:"Description,omitempty"`
	// A hash value for the externally stored object.
	DocumentHash []TextType `json:"DocumentHash,omitempty"`
	// A code signifying the encoding/decoding algorithm used with the external object.
	EncodingCode []CodeType `json:"EncodingCode,omitempty"`
	// The date on which availability of the resource can no longer be relied upon.
	ExpiryDate []DateType `json:"ExpiryDate,omitempty"`
	// The time after which availability of the resource can no longer be relied upon.
	ExpiryTime []TimeType `json:"ExpiryTime,omitempty"`
	// The file name of the external object.
	FileName []TextType `json:"FileName,omitempty"`
	// A code signifying the format of the external object.
	FormatCode []CodeType `json:"FormatCode,omitempty"`
	// A hash algorithm used to calculate the hash value of the externally stored object.
	HashAlgorithmMethod []TextType `json:"HashAlgorithmMethod,omitempty"`
	// A code signifying the mime type of the external object.
	MIMECode []CodeType `json:"MimeCode,omitempty"`
	// The Uniform Resource Identifier (URI) that identifies the external object as an Internet
	// resource.
	URI []IdentifierType `json:"URI,omitempty"`
}

// A list of two mutually exclusive Boolean values that express the only possible states of
// a property.
type IndicatorType struct {
	Empty bool `json:"_"`
}

// A class to describe a period of time.
type PeriodDetails struct {
	// A description of this period, expressed as text.
	Description []TextType `json:"Description,omitempty"`
	// A description of this period, expressed as a code.
	DescriptionCode []CodeType `json:"DescriptionCode,omitempty"`
	// The duration of this period, expressed as an ISO 8601 code.
	DurationMeasure []MeasureType `json:"DurationMeasure,omitempty"`
	// The date on which this period ends.
	EndDate []DateType `json:"EndDate,omitempty"`
	// The time at which this period ends.
	EndTime []TimeType `json:"EndTime,omitempty"`
	// The date on which this period begins.
	StartDate []DateType `json:"StartDate,omitempty"`
	// The time at which this period begins.
	StartTime []TimeType `json:"StartTime,omitempty"`
}

// A numeric value determined by measuring an object using a specified unit of measure.
type MeasureType struct {
	Empty                 float64 `json:"_"`
	UnitCode              string  `json:"unitCode"`
	UnitCodeListVersionID *string `json:"unitCodeListVersionID,omitempty"`
}

// A class to describe a financial account.
type FinancialAccountDetails struct {
	// A code signifying the format of this financial account.
	AccountFormatCode []CodeType `json:"AccountFormatCode,omitempty"`
	// A code signifying the type of this financial account.
	AccountTypeCode []CodeType `json:"AccountTypeCode,omitempty"`
	// An alias for the name of this financial account, to be used in place of the actual
	// account name for security reasons.
	AliasName []TextType `json:"AliasName,omitempty"`
	// The country in which the holder of the financial account is domiciled.
	Country []CountryDetails `json:"Country,omitempty"`
	// A code signifying the currency in which this financial account is held.
	CurrencyCode []CodeType `json:"CurrencyCode,omitempty"`
	// The branch of the financial institution associated with this financial account.
	FinancialInstitutionBranch []BranchDetails `json:"FinancialInstitutionBranch,omitempty"`
	// The identifier for this financial account; the bank account number.
	ID []IdentifierType `json:"ID,omitempty"`
	// The name of this financial account.
	Name []TextType `json:"Name,omitempty"`
	// Free-form text applying to the Payment for the owner of this account.
	PaymentNote []TextType `json:"PaymentNote,omitempty"`
}

// A class to describe a country.
type CountryDetails struct {
	// A code signifying this country.
	IdentificationCode []CodeType `json:"IdentificationCode,omitempty"`
	// The name of this country.
	Name []TextType `json:"Name,omitempty"`
}

// A class to describe a branch or a division of an organization.
type BranchDetails struct {
	// The address of this branch or division.
	Address []AddressDetails `json:"Address,omitempty"`
	// The financial institution that this branch belongs to (if applicable).
	FinancialInstitution []FinancialInstitutionDetails `json:"FinancialInstitution,omitempty"`
	// An identifier for this branch or division of an organization.
	ID []IdentifierType `json:"ID,omitempty"`
	// The name of this branch or division of an organization.
	Name []TextType `json:"Name,omitempty"`
}

// A class to define common information related to an address.
type AddressDetails struct {
	// An additional street name used to further clarify the address.
	AdditionalStreetName []TextType `json:"AdditionalStreetName,omitempty"`
	// A mutually agreed code signifying the format of this address.
	AddressFormatCode []CodeType `json:"AddressFormatCode,omitempty"`
	// An unstructured address line.
	AddressLine []AddressLineDetails `json:"AddressLine,omitempty"`
	// A mutually agreed code signifying the type of this address.
	AddressTypeCode []CodeType `json:"AddressTypeCode,omitempty"`
	// The name of the block (an area surrounded by streets and usually containing several
	// buildings) in which this address is located.
	BlockName []TextType `json:"BlockName,omitempty"`
	// The name of a building.
	BuildingName []TextType `json:"BuildingName,omitempty"`
	// The number of a building within the street.
	BuildingNumber []TextType `json:"BuildingNumber,omitempty"`
	// The name of a city, town, or village.
	CityName []TextType `json:"CityName,omitempty"`
	// The name of the subdivision of a city, town, or village in which this address is located,
	// such as the name of its district or borough.
	CitySubdivisionName []TextType `json:"CitySubdivisionName,omitempty"`
	// The country in which this address is situated.
	Country []CountryDetails `json:"Country,omitempty"`
	// The political or administrative division of a country in which this address is located,
	// such as the name of its county, province, or state, expressed as text.
	CountrySubentity []TextType `json:"CountrySubentity,omitempty"`
	// The political or administrative division of a country in which this address is located,
	// such as a county, province, or state, expressed as a code (typically nationally agreed).
	CountrySubentityCode []CodeType `json:"CountrySubentityCode,omitempty"`
	// The department of the addressee.
	Department []TextType `json:"Department,omitempty"`
	// The district or geographical division of a country or region in which this address is
	// located.
	District []TextType `json:"District,omitempty"`
	// An identifiable floor of a building.
	Floor []TextType `json:"Floor,omitempty"`
	// An identifier for this address within an agreed scheme of address identifiers.
	ID []IdentifierType `json:"ID,omitempty"`
	// The specific identifable location within a building where mail is delivered.
	InhouseMail []TextType `json:"InhouseMail,omitempty"`
	// The geographical coordinates of this address.
	LocationCoordinate []LocationCoordinateDetails `json:"LocationCoordinate,omitempty"`
	// The name, expressed as text, of a person or department in an organization to whose
	// attention incoming mail is directed; corresponds to the printed forms "for the attention
	// of", "FAO", and ATTN:".
	MarkAttention []TextType `json:"MarkAttention,omitempty"`
	// The name, expressed as text, of a person or organization at this address into whose care
	// incoming mail is entrusted; corresponds to the printed forms "care of" and "c/o".
	MarkCare []TextType `json:"MarkCare,omitempty"`
	// An identifier (e.g., a parcel number) for the piece of land associated with this address.
	PlotIdentification []TextType `json:"PlotIdentification,omitempty"`
	// The postal identifier for this address according to the relevant national postal service,
	// such as a ZIP code or Post Code.
	PostalZone []TextType `json:"PostalZone,omitempty"`
	// A post office box number registered for postal delivery by a postal service provider.
	Postbox []TextType `json:"Postbox,omitempty"`
	// The recognized geographic or economic region or group of countries in which this address
	// is located.
	Region []TextType `json:"Region,omitempty"`
	// An identifiable room, suite, or apartment of a building.
	Room []TextType `json:"Room,omitempty"`
	// The name of the street, road, avenue, way, etc. to which the number of the building is
	// attached.
	StreetName []TextType `json:"StreetName,omitempty"`
	// The time zone in which this address is located (as an offset from Universal Coordinated
	// Time (UTC)) at the time of exchange.
	TimezoneOffset []TextType `json:"TimezoneOffset,omitempty"`
}

// A class to define an unstructured address line.
type AddressLineDetails struct {
	// An address line expressed as unstructured text.
	Line []TextType `json:"Line"`
}

// A class for defining a set of geographical coordinates (apparently misnamed).
type LocationCoordinateDetails struct {
	// The altitude of the location.
	AltitudeMeasure []MeasureType `json:"AltitudeMeasure,omitempty"`
	// A code signifying the location system used.
	CoordinateSystemCode []CodeType `json:"CoordinateSystemCode,omitempty"`
	// The degree component of a latitude measured in degrees and minutes.
	LatitudeDegreesMeasure []MeasureType `json:"LatitudeDegreesMeasure,omitempty"`
	// A code signifying the direction of latitude measurement from the equator (north or south).
	LatitudeDirectionCode []CodeType `json:"LatitudeDirectionCode,omitempty"`
	// The minutes component of a latitude measured in degrees and minutes (modulo 60).
	LatitudeMinutesMeasure []MeasureType `json:"LatitudeMinutesMeasure,omitempty"`
	// The degree component of a longitude measured in degrees and minutes.
	LongitudeDegreesMeasure []MeasureType `json:"LongitudeDegreesMeasure,omitempty"`
	// A code signifying the direction of longitude measurement from the prime meridian (east or
	// west).
	LongitudeDirectionCode []CodeType `json:"LongitudeDirectionCode,omitempty"`
	// The minutes component of a longitude measured in degrees and minutes (modulo 60).
	LongitudeMinutesMeasure []MeasureType `json:"LongitudeMinutesMeasure,omitempty"`
}

// A class to describe a financial institution.
type FinancialInstitutionDetails struct {
	// The address of this financial institution.
	Address []AddressDetails `json:"Address,omitempty"`
	// An identifier for this financial institution. It is recommended that the ISO 9362 Bank
	// Identification Code (BIC) be used as the ID.
	ID []IdentifierType `json:"ID,omitempty"`
	// The name of this financial institution.
	Name []TextType `json:"Name,omitempty"`
}

// Numeric information that is assigned or is determined by calculation, counting, or
// sequencing and is expressed as a percentage. It does not require a unit of quantity or
// unit of measure.
//
// Numeric information that is assigned or is determined by calculation, counting, or
// sequencing. It does not require a unit of quantity or unit of measure.
//
// A numeric expression of a rate that is assigned or is determined by calculation,
// counting, or sequencing. It does not require a unit of quantity or unit of measure.
type NumericType struct {
	Empty  float64 `json:"_"`
	Format *string `json:"format,omitempty"`
}

// A class to describe a scheme for corporate registration.
type CorporateRegistrationSchemeDetails struct {
	// A code signifying the type of this registration scheme.
	CorporateRegistrationTypeCode []CodeType `json:"CorporateRegistrationTypeCode,omitempty"`
	// An identifier for this registration scheme.
	ID []IdentifierType `json:"ID,omitempty"`
	// A geographic area in which this registration scheme applies.
	JurisdictionRegionAddress []AddressDetails `json:"JurisdictionRegionAddress,omitempty"`
	// The name of this registration scheme.
	Name []TextType `json:"Name,omitempty"`
}

// A number of monetary units specified using a given unit of currency.
type AmountType struct {
	Empty                     float64 `json:"_"`
	CurrencyCodeListVersionID *string `json:"currencyCodeListVersionID,omitempty"`
	CurrencyID                string  `json:"currencyID"`
}

// A class to describe a language.
type LanguageDetails struct {
	// An identifier for this language.
	ID []IdentifierType `json:"ID,omitempty"`
	// A code signifying the locale in which this language is used.
	LocaleCode []CodeType `json:"LocaleCode,omitempty"`
	// The name of this language.
	Name []TextType `json:"Name,omitempty"`
}

// A class to define an identifier for a party.
type PartyIdentificationDetails struct {
	// An identifier for the party.
	ID []IdentifierType `json:"ID"`
}

// A class for defining the name of a party.
type PartyNameDetails struct {
	// The name of the party.
	Name []TextType `json:"Name"`
}

// A class to describe a taxation scheme applying to a party.
type PartyTaxSchemeDetails struct {
	// An identifier for the party assigned for tax purposes by the taxation authority.
	CompanyID []IdentifierType `json:"CompanyID,omitempty"`
	// A reason for the party's exemption from tax, expressed as text.
	ExemptionReason []TextType `json:"ExemptionReason,omitempty"`
	// A reason for the party's exemption from tax, expressed as a code.
	ExemptionReasonCode []CodeType `json:"ExemptionReasonCode,omitempty"`
	// The address of the party as registered for tax purposes.
	RegistrationAddress []AddressDetails `json:"RegistrationAddress,omitempty"`
	// The name of the party as registered with the relevant fiscal authority.
	RegistrationName []TextType `json:"RegistrationName,omitempty"`
	// A code signifying the tax level applicable to the party within this taxation scheme.
	TaxLevelCode []CodeType `json:"TaxLevelCode,omitempty"`
	// The taxation scheme applicable to the party.
	TaxScheme []TaxSchemeDetails `json:"TaxScheme"`
}

// A class to describe a taxation scheme (e.g., VAT, State tax, County tax).
type TaxSchemeDetails struct {
	// A code signifying the currency in which the tax is collected and reported.
	CurrencyCode []CodeType `json:"CurrencyCode,omitempty"`
	// An identifier for this taxation scheme.
	ID []IdentifierType `json:"ID,omitempty"`
	// A geographic area in which this taxation scheme applies.
	JurisdictionRegionAddress []AddressDetails `json:"JurisdictionRegionAddress,omitempty"`
	// The name of this taxation scheme.
	Name []TextType `json:"Name,omitempty"`
	// A code signifying the type of tax.
	TaxTypeCode []CodeType `json:"TaxTypeCode,omitempty"`
}

// A class to describe a location.
type LocationDetails struct {
	// The address of this location.
	Address []AddressDetails `json:"Address,omitempty"`
	// Free-form text describing the physical conditions of the location.
	Conditions []TextType `json:"Conditions,omitempty"`
	// A territorial division of a country, such as a county or state, expressed as text.
	CountrySubentity []TextType `json:"CountrySubentity,omitempty"`
	// A territorial division of a country, such as a county or state, expressed as a code.
	CountrySubentityCode []CodeType `json:"CountrySubentityCode,omitempty"`
	// Text describing this location.
	Description []TextType `json:"Description,omitempty"`
	// An identifier for this location, e.g., the EAN Location Number, GLN.
	ID []IdentifierType `json:"ID,omitempty"`
	// The Uniform Resource Identifier (URI) of a document providing information about this
	// location.
	InformationURI []IdentifierType `json:"InformationURI,omitempty"`
	// The geographical coordinates of this location.
	LocationCoordinate []LocationCoordinateDetails `json:"LocationCoordinate,omitempty"`
	// A code signifying the type of location.
	LocationTypeCode []CodeType `json:"LocationTypeCode,omitempty"`
	// The name of this location.
	Name []TextType `json:"Name,omitempty"`
	// A location subsidiary to this location.
	SubsidiaryLocation []LocationDetails `json:"SubsidiaryLocation,omitempty"`
	// A period during which this location can be used (e.g., for delivery).
	ValidityPeriod []PeriodDetails `json:"ValidityPeriod,omitempty"`
}

// A class to describe a supplier party.
type SupplierPartyDetails struct {
	// A contact at this supplier party for accounting.
	AccountingContact []ContactDetails `json:"AccountingContact,omitempty"`
	// An additional identifier for this supplier party.
	AdditionalAccountID []IdentifierType `json:"AdditionalAccountID,omitempty"`
	// An identifier for this supplier party, assigned by the customer.
	CustomerAssignedAccountID []IdentifierType `json:"CustomerAssignedAccountID,omitempty"`
	// Text describing the supplier's ability to send invoice data via a purchase card provider
	// (e.g., VISA, MasterCard, American Express).
	DataSendingCapability []TextType `json:"DataSendingCapability,omitempty"`
	// A contact at this supplier party for despatches (pickups).
	DespatchContact []ContactDetails `json:"DespatchContact,omitempty"`
	// The supplier party itself.
	Party []PartyDetails `json:"Party,omitempty"`
	// The primary contact for this supplier party.
	SellerContact []ContactDetails `json:"SellerContact,omitempty"`
}

// A class to describe information about a charge or discount as applied to a price
// component.
type AllowanceChargeDetails struct {
	// The accounting cost centre used by the buyer to account for this allowance or charge,
	// expressed as text.
	AccountingCost []TextType `json:"AccountingCost,omitempty"`
	// The accounting cost centre used by the buyer to account for this allowance or charge,
	// expressed as a code.
	AccountingCostCode []CodeType `json:"AccountingCostCode,omitempty"`
	// The reason for this allowance or charge.
	AllowanceChargeReason []TextType `json:"AllowanceChargeReason,omitempty"`
	// A mutually agreed code signifying the reason for this allowance or charge.
	AllowanceChargeReasonCode []CodeType `json:"AllowanceChargeReasonCode,omitempty"`
	// The monetary amount of this allowance or charge to be applied.
	Amount []AmountType `json:"Amount"`
	// The monetary amount to which the multiplier factor is applied in calculating the amount
	// of this allowance or charge.
	BaseAmount []AmountType `json:"BaseAmount,omitempty"`
	// An indicator that this AllowanceCharge describes a charge (true) or a discount (false).
	ChargeIndicator []IndicatorType `json:"ChargeIndicator"`
	// An identifier for this allowance or charge.
	ID []IdentifierType `json:"ID,omitempty"`
	// A number by which the base amount is multiplied to calculate the actual amount of this
	// allowance or charge.
	MultiplierFactorNumeric []NumericType `json:"MultiplierFactorNumeric,omitempty"`
	// A means of payment for this allowance or charge.
	PaymentMeans []PaymentMeansDetails `json:"PaymentMeans,omitempty"`
	// The allowance or charge per item; the total allowance or charge is calculated by
	// multiplying the per unit amount by the quantity of items, either at the level of the
	// individual transaction line or for the total number of items in the document, depending
	// on the context in which it appears.
	PerUnitAmount []AmountType `json:"PerUnitAmount,omitempty"`
	// An indicator that this allowance or charge is prepaid (true) or not (false).
	PrepaidIndicator []IndicatorType `json:"PrepaidIndicator,omitempty"`
	// A number indicating the order of this allowance or charge in the sequence of calculations
	// applied when there are multiple allowances or charges.
	SequenceNumeric []NumericType `json:"SequenceNumeric,omitempty"`
	// A tax category applicable to this allowance or charge.
	TaxCategory []TaxCategoryDetails `json:"TaxCategory,omitempty"`
	// The total of all the taxes applicable to this allowance or charge.
	TaxTotal []TaxTotalDetails `json:"TaxTotal,omitempty"`
}

// A class to describe a means of payment.
type PaymentMeansDetails struct {
	// A credit card, debit card, or charge card account that constitutes this means of payment.
	CardAccount []CardAccountDetails `json:"CardAccount,omitempty"`
	// A credit account associated with this means of payment.
	CreditAccount []CreditAccountDetails `json:"CreditAccount,omitempty"`
	// An identifier for this means of payment.
	ID []IdentifierType `json:"ID,omitempty"`
	// An identifier for the payment instruction.
	InstructionID []IdentifierType `json:"InstructionID,omitempty"`
	// Free-form text conveying information that is not contained explicitly in other structures.
	InstructionNote []TextType `json:"InstructionNote,omitempty"`
	// The payee's financial account.
	PayeeFinancialAccount []FinancialAccountDetails `json:"PayeeFinancialAccount,omitempty"`
	// The payer's financial account.
	PayerFinancialAccount []FinancialAccountDetails `json:"PayerFinancialAccount,omitempty"`
	// A code signifying the payment channel for this means of payment.
	PaymentChannelCode []CodeType `json:"PaymentChannelCode,omitempty"`
	// The date on which payment is due for this means of payment.
	PaymentDueDate []DateType `json:"PaymentDueDate,omitempty"`
	// An identifier for a payment made using this means of payment.
	PaymentID []IdentifierType `json:"PaymentID,omitempty"`
	// The payment mandate associated with this means of payment.
	PaymentMandate []PaymentMandateDetails `json:"PaymentMandate,omitempty"`
	// A code signifying the type of this means of payment.
	PaymentMeansCode []CodeType `json:"PaymentMeansCode"`
	// A trade finance agreement applicable to this means of payment.
	TradeFinancing []TradeFinancingDetails `json:"TradeFinancing,omitempty"`
}

// A class to define a credit card, debit card, or charge card account.
type CardAccountDetails struct {
	// A mutually agreed code to distinguish between CHIP and MAG STRIPE cards.
	CardChipCode []CodeType `json:"CardChipCode,omitempty"`
	// A mutually agreed code signifying the type of card. Examples of types are "debit",
	// "credit" and "purchasing"
	CardTypeCode []CodeType `json:"CardTypeCode,omitempty"`
	// An identifier on the chip card for the application that provides the quoted information;
	// an AID (application ID).
	ChipApplicationID []IdentifierType `json:"ChipApplicationID,omitempty"`
	// An identifier for the Card Verification Value (often found on the reverse of the card
	// itself).
	Cv2ID []IdentifierType `json:"CV2ID,omitempty"`
	// The date on which the card expires.
	ExpiryDate []DateType `json:"ExpiryDate,omitempty"`
	// The name of the cardholder.
	HolderName []TextType `json:"HolderName,omitempty"`
	// An identifier for the card, specified by the issuer.
	IssueNumberID []IdentifierType `json:"IssueNumberID,omitempty"`
	// An identifier for the institution issuing the card.
	IssuerID []IdentifierType `json:"IssuerID,omitempty"`
	// An identifier for the financial service network provider of the card.
	NetworkID []IdentifierType `json:"NetworkID"`
	// An identifier of the card (e.g., the Primary Account Number (PAN)).
	PrimaryAccountNumberID []IdentifierType `json:"PrimaryAccountNumberID"`
	// The date from which the card is valid.
	ValidityStartDate []DateType `json:"ValidityStartDate,omitempty"`
}

// A class to identify a credit account for sales on account.
type CreditAccountDetails struct {
	// An identifier for this credit account.
	AccountID []IdentifierType `json:"AccountID"`
}

// A class to describe a payment mandate.
type PaymentMandateDetails struct {
	// A clause applicable to this payment mandate.
	Clause []ClauseDetails `json:"Clause,omitempty"`
	// An identifier for this payment mandate.
	ID []IdentifierType `json:"ID,omitempty"`
	// A code signifying the type of this payment mandate.
	MandateTypeCode []CodeType `json:"MandateTypeCode,omitempty"`
	// The maximum amount to be paid within a single instruction.
	MaximumPaidAmount []AmountType `json:"MaximumPaidAmount,omitempty"`
	// The number of maximum payment instructions allowed within the validity period.
	MaximumPaymentInstructionsNumeric []NumericType `json:"MaximumPaymentInstructionsNumeric,omitempty"`
	// The payer's financial account.
	PayerFinancialAccount []FinancialAccountDetails `json:"PayerFinancialAccount,omitempty"`
	// The payer party (if different from the debtor).
	PayerParty []PartyDetails `json:"PayerParty,omitempty"`
	// The period of the reverse payment.
	PaymentReversalPeriod []PeriodDetails `json:"PaymentReversalPeriod,omitempty"`
	// An identifier for a signature applied by a signatory party.
	SignatureID []IdentifierType `json:"SignatureID,omitempty"`
	// The period during which this mandate is valid.
	ValidityPeriod []PeriodDetails `json:"ValidityPeriod,omitempty"`
}

// A class to define a clause (a distinct article or provision) in a contract, treaty, will,
// or other formal or legal written document requiring compliance.
type ClauseDetails struct {
	// The text of this clause.
	Content []TextType `json:"Content,omitempty"`
	// An identifier for this clause.
	ID []IdentifierType `json:"ID,omitempty"`
}

// A class to describe a trade financing instrument.
type TradeFinancingDetails struct {
	// A clause applicable to this trade financing instrument.
	Clause []ClauseDetails `json:"Clause,omitempty"`
	// A reference to a contract document.
	ContractDocumentReference []DocumentReferenceDetails `json:"ContractDocumentReference,omitempty"`
	// A reference to a document associated with this trade financing instrument.
	DocumentReference []DocumentReferenceDetails `json:"DocumentReference,omitempty"`
	// An internal bank account used by the bank or its first agent to manage the line of credit
	// granted to the financing requester.
	FinancingFinancialAccount []FinancialAccountDetails `json:"FinancingFinancialAccount,omitempty"`
	// A code signifying the type of this financing instrument.
	FinancingInstrumentCode []CodeType `json:"FinancingInstrumentCode,omitempty"`
	// The financing party (bank or other enabled party).
	FinancingParty []PartyDetails `json:"FinancingParty"`
	// An identifier for this trade financing instrument.
	ID []IdentifierType `json:"ID,omitempty"`
}

// A class to describe one of the tax categories within a taxation scheme (e.g., High Rate
// VAT, Low Rate VAT).
type TaxCategoryDetails struct {
	// A Unit of Measures used as the basic for the tax calculation applied at a certain rate
	// per unit.
	BaseUnitMeasure []MeasureType `json:"BaseUnitMeasure,omitempty"`
	// An identifier for this tax category.
	ID []IdentifierType `json:"ID,omitempty"`
	// The name of this tax category.
	Name []TextType `json:"Name,omitempty"`
	// The tax rate for this category, expressed as a percentage.
	Percent []NumericType `json:"Percent,omitempty"`
	// Where a tax is applied at a certain rate per unit, the rate per unit applied.
	PerUnitAmount []AmountType `json:"PerUnitAmount,omitempty"`
	// The reason for tax being exempted, expressed as text.
	TaxExemptionReason []TextType `json:"TaxExemptionReason,omitempty"`
	// The reason for tax being exempted, expressed as a code.
	TaxExemptionReasonCode []CodeType `json:"TaxExemptionReasonCode,omitempty"`
	// The taxation scheme within which this tax category is defined.
	TaxScheme []TaxSchemeDetails `json:"TaxScheme"`
	// Where a tax is tiered, the range of taxable amounts that determines the rate of tax
	// applicable to this tax category.
	TierRange []TextType `json:"TierRange,omitempty"`
	// Where a tax is tiered, the tax rate that applies within the specified range of taxable
	// amounts for this tax category.
	TierRatePercent []NumericType `json:"TierRatePercent,omitempty"`
}

// A class to describe the total tax for a particular taxation scheme.
type TaxTotalDetails struct {
	// The rounding amount (positive or negative) added to the calculated tax total to produce
	// the rounded TaxAmount.
	RoundingAmount []AmountType `json:"RoundingAmount,omitempty"`
	// The total tax amount for a particular taxation scheme, e.g., VAT; the sum of the tax
	// subtotals for each tax category within the taxation scheme.
	TaxAmount []AmountType `json:"TaxAmount"`
	// An indicator that this total is recognized as legal evidence for taxation purposes (true)
	// or not (false).
	TaxEvidenceIndicator []IndicatorType `json:"TaxEvidenceIndicator,omitempty"`
	// An indicator that tax is included in the calculation (true) or not (false).
	TaxIncludedIndicator []IndicatorType `json:"TaxIncludedIndicator,omitempty"`
	// One of the subtotals the sum of which equals the total tax amount for a particular
	// taxation scheme.
	TaxSubtotal []TaxSubtotalDetails `json:"TaxSubtotal,omitempty"`
}

// A class to define the subtotal for a particular tax category within a particular taxation
// scheme, such as standard rate within VAT.
type TaxSubtotalDetails struct {
	// The unit of measure on which the tax calculation is based
	BaseUnitMeasure []MeasureType `json:"BaseUnitMeasure,omitempty"`
	// The number of this tax subtotal in the sequence of subtotals corresponding to the order
	// in which multiple taxes are applied. If all taxes are applied to the same taxable amount
	// (i.e., their order of application is inconsequential), then CalculationSequenceNumeric is
	// 1 for all tax subtotals applied to a given amount.
	CalculationSequenceNumeric []NumericType `json:"CalculationSequenceNumeric,omitempty"`
	// The tax rate of the tax category applied to this tax subtotal, expressed as a percentage.
	Percent []NumericType `json:"Percent,omitempty"`
	// Where a tax is applied at a certain rate per unit, the rate per unit applied.
	PerUnitAmount []AmountType `json:"PerUnitAmount,omitempty"`
	// The net amount to which the tax percent (rate) is applied to calculate the tax amount.
	TaxableAmount []AmountType `json:"TaxableAmount,omitempty"`
	// The amount of this tax subtotal.
	TaxAmount []AmountType `json:"TaxAmount"`
	// The tax category applicable to this subtotal.
	TaxCategory []TaxCategoryDetails `json:"TaxCategory"`
	// Where a tax is tiered, the range of taxable amounts that determines the rate of tax
	// applicable to this tax subtotal.
	TierRange []TextType `json:"TierRange,omitempty"`
	// Where a tax is tiered, the tax rate that applies within a specified range of taxable
	// amounts for this tax subtotal.
	TierRatePercent []NumericType `json:"TierRatePercent,omitempty"`
	// The amount of this tax subtotal, expressed in the currency used for invoicing.
	TransactionCurrencyTaxAmount []AmountType `json:"TransactionCurrencyTaxAmount,omitempty"`
}

// A class to define a reference to a billing document.
type BillingReferenceDetails struct {
	// A reference to an additional document.
	AdditionalDocumentReference []DocumentReferenceDetails `json:"AdditionalDocumentReference,omitempty"`
	// A reference to a transaction line in the billing document.
	BillingReferenceLine []BillingReferenceLineDetails `json:"BillingReferenceLine,omitempty"`
	// A reference to a credit note.
	CreditNoteDocumentReference []DocumentReferenceDetails `json:"CreditNoteDocumentReference,omitempty"`
	// A reference to a debit note.
	DebitNoteDocumentReference []DocumentReferenceDetails `json:"DebitNoteDocumentReference,omitempty"`
	// A reference to an invoice.
	InvoiceDocumentReference []DocumentReferenceDetails `json:"InvoiceDocumentReference,omitempty"`
	// A reference to a billing reminder.
	ReminderDocumentReference []DocumentReferenceDetails `json:"ReminderDocumentReference,omitempty"`
	// A reference to a self billed credit note.
	SelfBilledCreditNoteDocumentReference []DocumentReferenceDetails `json:"SelfBilledCreditNoteDocumentReference,omitempty"`
	// A reference to a self billed invoice.
	SelfBilledInvoiceDocumentReference []DocumentReferenceDetails `json:"SelfBilledInvoiceDocumentReference,omitempty"`
}

// A class to define a reference to a transaction line in a billing document.
type BillingReferenceLineDetails struct {
	// An allowance or charge applicable to the transaction line.
	AllowanceCharge []AllowanceChargeDetails `json:"AllowanceCharge,omitempty"`
	// The monetary amount of the transaction line, including any allowances and charges but
	// excluding taxes.
	Amount []AmountType `json:"Amount,omitempty"`
	// An identifier for this transaction line in a billing document.
	ID []IdentifierType `json:"ID"`
}

// A class to define a line in a Credit Note or Self Billed Credit Note.
type CreditNoteLineDetails struct {
	// The buyer's accounting cost centre for this credit note line, expressed as text.
	AccountingCost []TextType `json:"AccountingCost,omitempty"`
	// The buyer's accounting cost centre for this credit note line, expressed as a code.
	AccountingCostCode []CodeType `json:"AccountingCostCode,omitempty"`
	// An allowance or charge associated with this credit note.
	AllowanceCharge []AllowanceChargeDetails `json:"AllowanceCharge,omitempty"`
	// A reference to a billing document associated with this credit note line.
	BillingReference []BillingReferenceDetails `json:"BillingReference,omitempty"`
	// The quantity of items credited in this credit note line.
	CreditedQuantity []QuantityType `json:"CreditedQuantity,omitempty"`
	// A delivery associated with this credit note line.
	Delivery []DeliveryDetails `json:"Delivery,omitempty"`
	// Terms and conditions of a delivery associated with this credit note line.
	DeliveryTerms []DeliveryTermsDetails `json:"DeliveryTerms,omitempty"`
	// A reference to a despatch line associated with this credit note line.
	DespatchLineReference []LineReferenceDetails `json:"DespatchLineReference,omitempty"`
	// A reason for the credit.
	DiscrepancyResponse []ResponseDetails `json:"DiscrepancyResponse,omitempty"`
	// A reference to a document associated with this credit note line.
	DocumentReference []DocumentReferenceDetails `json:"DocumentReference,omitempty"`
	// An indicator that this credit note line is free of charge (true) or not (false). The
	// default is false.
	FreeOfChargeIndicator []IndicatorType `json:"FreeOfChargeIndicator,omitempty"`
	// An identifier for this credit note line.
	ID []IdentifierType `json:"ID"`
	// An invoice period to which this credit note line applies.
	InvoicePeriod []PeriodDetails `json:"InvoicePeriod,omitempty"`
	// The item associated with this credit note line.
	Item []ItemDetails `json:"Item,omitempty"`
	// The price extension, calculated by multiplying the price per unit by the quantity of
	// items on this credit note line.
	ItemPriceExtension []PriceExtensionDetails `json:"ItemPriceExtension,omitempty"`
	// The total amount for this credit note line, including allowance charges but exclusive of
	// taxes.
	LineExtensionAmount []AmountType `json:"LineExtensionAmount,omitempty"`
	// Free-form text conveying information that is not contained explicitly in other structures.
	Note []TextType `json:"Note,omitempty"`
	// A reference to an order line associated with this credit note line.
	OrderLineReference []OrderLineReferenceDetails `json:"OrderLineReference,omitempty"`
	// The party who originated the Order to which the Credit Note is related.
	OriginatorParty []PartyDetails `json:"OriginatorParty,omitempty"`
	// A code signifying the business purpose for this payment.
	PaymentPurposeCode []CodeType `json:"PaymentPurposeCode,omitempty"`
	// A specification of payment terms associated with this credit note line.
	PaymentTerms []PaymentTermsDetails `json:"PaymentTerms,omitempty"`
	// The price of the item associated with this credit note line.
	Price []PriceDetails `json:"Price,omitempty"`
	// A reference to pricing and item location information associated with this credit note
	// line.
	PricingReference []PricingReferenceDetails `json:"PricingReference,omitempty"`
	// A reference to a receipt line associated with this credit note line.
	ReceiptLineReference []LineReferenceDetails `json:"ReceiptLineReference,omitempty"`
	// A class defining one or more Credit Note Lines detailing the credit note line.
	SubCreditNoteLine []CreditNoteLineDetails `json:"SubCreditNoteLine,omitempty"`
	// The date of this credit note line, used to indicate the point at which tax becomes
	// applicable.
	TaxPointDate []DateType `json:"TaxPointDate,omitempty"`
	// A total amount of taxes of a particular kind applicable to this credit note line.
	TaxTotal []TaxTotalDetails `json:"TaxTotal,omitempty"`
	// A universally unique identifier for this credit note line.
	UUID []IdentifierType `json:"UUID,omitempty"`
}

// A counted number of non-monetary units, possibly including a fractional part.
//
// A counted number of non-monetary units possibly including fractions.
type QuantityType struct {
	Empty                  float64 `json:"_"`
	UnitCode               *string `json:"unitCode,omitempty"`
	UnitCodeListAgencyID   *string `json:"unitCodeListAgencyID,omitempty"`
	UnitCodeListAgencyName *string `json:"unitCodeListAgencyName,omitempty"`
	UnitCodeListID         *string `json:"unitCodeListID,omitempty"`
}

// A class to define a line in a Receipt Advice.
type ReceiptLineDetails struct {
	// A reference to a despatch line associated with this receipt line.
	DespatchLineReference []LineReferenceDetails `json:"DespatchLineReference,omitempty"`
	// A reference to a document associated with this receipt line.
	DocumentReference []DocumentReferenceDetails `json:"DocumentReference,omitempty"`
	// An identifier for this receipt line.
	ID []IdentifierType `json:"ID"`
	// An item associated with this receipt line.
	Item []ItemDetails `json:"Item,omitempty"`
	// Free-form text conveying information that is not contained explicitly in other structures.
	Note []TextType `json:"Note,omitempty"`
	// A reference to the order line associated with this receipt line.
	OrderLineReference []OrderLineReferenceDetails `json:"OrderLineReference,omitempty"`
	// The quantity over-supplied, i.e., the quantity over and above the quantity ordered.
	OversupplyQuantity []QuantityType `json:"OversupplyQuantity,omitempty"`
	// A code signifying the type of a discrepancy in quantity.
	QuantityDiscrepancyCode []CodeType `json:"QuantityDiscrepancyCode,omitempty"`
	// The date on which the goods or services were received.
	ReceivedDate []DateType `json:"ReceivedDate,omitempty"`
	// The quantity received.
	ReceivedQuantity []QuantityType `json:"ReceivedQuantity,omitempty"`
	// A code signifying the action that the delivery party wishes the despatch party to take as
	// the result of a rejection.
	RejectActionCode []CodeType `json:"RejectActionCode,omitempty"`
	// The quantity rejected.
	RejectedQuantity []QuantityType `json:"RejectedQuantity,omitempty"`
	// The reason for a rejection, expressed as text.
	RejectReason []TextType `json:"RejectReason,omitempty"`
	// The reason for a rejection, expressed as a code.
	RejectReasonCode []CodeType `json:"RejectReasonCode,omitempty"`
	// A shipment associated with this receipt line.
	Shipment []ShipmentDetails `json:"Shipment,omitempty"`
	// A code signifying the action that the delivery party wishes the despatch party to take as
	// the result of a shortage.
	ShortageActionCode []CodeType `json:"ShortageActionCode,omitempty"`
	// The quantity received short; the difference between the quantity reported despatched and
	// the quantity actually received.
	ShortQuantity []QuantityType `json:"ShortQuantity,omitempty"`
	// A complaint about the timing of delivery, expressed as text.
	TimingComplaint []TextType `json:"TimingComplaint,omitempty"`
	// A complaint about the timing of delivery, expressed as a code.
	TimingComplaintCode []CodeType `json:"TimingComplaintCode,omitempty"`
	// A universally unique identifier for this receipt line.
	UUID []IdentifierType `json:"UUID,omitempty"`
}

// A class to define a line in a Despatch Advice.
type DespatchLineDetails struct {
	// The quantity on back order at the supplier.
	BackorderQuantity []QuantityType `json:"BackorderQuantity,omitempty"`
	// The reason for the back order.
	BackorderReason []TextType `json:"BackorderReason,omitempty"`
	// The quantity despatched (picked up).
	DeliveredQuantity []QuantityType `json:"DeliveredQuantity,omitempty"`
	// A reference to a document associated with this despatch line.
	DocumentReference []DocumentReferenceDetails `json:"DocumentReference,omitempty"`
	// An identifier for this despatch line.
	ID []IdentifierType `json:"ID"`
	// The item associated with this despatch line.
	Item []ItemDetails `json:"Item"`
	// A code signifying the status of this despatch line with respect to its original state.
	LineStatusCode []CodeType `json:"LineStatusCode,omitempty"`
	// Free-form text conveying information that is not contained explicitly in other structures.
	Note []TextType `json:"Note,omitempty"`
	// A reference to an order line associated with this despatch line.
	OrderLineReference []OrderLineReferenceDetails `json:"OrderLineReference"`
	// The quantity outstanding (which will follow in a later despatch).
	OutstandingQuantity []QuantityType `json:"OutstandingQuantity,omitempty"`
	// The reason for the outstanding quantity.
	OutstandingReason []TextType `json:"OutstandingReason,omitempty"`
	// The quantity over-supplied, i.e., the quantity over and above that ordered.
	OversupplyQuantity []QuantityType `json:"OversupplyQuantity,omitempty"`
	// A shipment associated with this despatch line.
	Shipment []ShipmentDetails `json:"Shipment,omitempty"`
	// A universally unique identifier for this despatch line.
	UUID []IdentifierType `json:"UUID,omitempty"`
}

// A class to describe a uniquely identifiable unit consisting of one or more packages,
// goods items, or pieces of transport equipment.
type TransportHandlingUnitDetails struct {
	// A package contained in this transport handling unit.
	ActualPackage []PackageDetails `json:"ActualPackage,omitempty"`
	// Describes identifiers or references relating to customs procedures.
	CustomsDeclaration []CustomsDeclarationDetails `json:"CustomsDeclaration,omitempty"`
	// Text describing damage associated with this transport handling unit.
	DamageRemarks []TextType `json:"DamageRemarks,omitempty"`
	// The floor space measurement dimension associated with this transport handling unit.
	FloorSpaceMeasurementDimension []DimensionDetails `json:"FloorSpaceMeasurementDimension,omitempty"`
	// A goods item contained in this transport handling unit.
	GoodsItem []GoodsItemDetails `json:"GoodsItem,omitempty"`
	// The handling required for this transport handling unit, expressed as a code.
	HandlingCode []CodeType `json:"HandlingCode,omitempty"`
	// The handling required for this transport handling unit, expressed as text.
	HandlingInstructions []TextType `json:"HandlingInstructions,omitempty"`
	// A despatch line associated with this transport handling unit.
	HandlingUnitDespatchLine []DespatchLineDetails `json:"HandlingUnitDespatchLine,omitempty"`
	// Transit-related information regarding a type of hazardous goods contained in this
	// transport handling unit.
	HazardousGoodsTransit []HazardousGoodsTransitDetails `json:"HazardousGoodsTransit,omitempty"`
	// An indicator that the materials contained in this transport handling unit are subject to
	// an international regulation concerning the carriage of dangerous goods (true) or not
	// (false).
	HazardousRiskIndicator []IndicatorType `json:"HazardousRiskIndicator,omitempty"`
	// An identifier for this transport handling unit.
	ID []IdentifierType `json:"ID,omitempty"`
	// The maximum allowable operating temperature of this transport handling unit.
	MaximumTemperature []TemperatureDetails `json:"MaximumTemperature,omitempty"`
	// A measurable dimension (length, mass, weight, or volume) of this transport handling unit.
	MeasurementDimension []DimensionDetails `json:"MeasurementDimension,omitempty"`
	// The minimum required operating temperature of this transport handling unit.
	MinimumTemperature []TemperatureDetails `json:"MinimumTemperature,omitempty"`
	// A package contained in this transport handling unit.
	Package []PackageDetails `json:"Package,omitempty"`
	// The pallet space measurement dimension associated to this transport handling unit.
	PalletSpaceMeasurementDimension []DimensionDetails `json:"PalletSpaceMeasurementDimension,omitempty"`
	// A receipt line associated with this transport handling unit.
	ReceivedHandlingUnitReceiptLine []ReceiptLineDetails `json:"ReceivedHandlingUnitReceiptLine,omitempty"`
	// A shipment associated with this transport handling unit.
	ReferencedShipment []ShipmentDetails `json:"ReferencedShipment,omitempty"`
	// A reference to a shipping document associated with this transport handling unit.
	ShipmentDocumentReference []DocumentReferenceDetails `json:"ShipmentDocumentReference,omitempty"`
	// Text describing the marks and numbers on this transport handling unit.
	ShippingMarks []TextType `json:"ShippingMarks,omitempty"`
	// The status of this transport handling unit.
	Status []StatusDetails `json:"Status,omitempty"`
	// The total number of goods items in this transport handling unit.
	TotalGoodsItemQuantity []QuantityType `json:"TotalGoodsItemQuantity,omitempty"`
	// The total number of packages in this transport handling unit.
	TotalPackageQuantity []QuantityType `json:"TotalPackageQuantity,omitempty"`
	// An identifier for use in tracing this transport handling unit, such as the EPC number
	// used in RFID.
	TraceID []IdentifierType `json:"TraceID,omitempty"`
	// A piece of transport equipment associated with this transport handling unit.
	TransportEquipment []TransportEquipmentDetails `json:"TransportEquipment,omitempty"`
	// A code signifying the type of this transport handling unit.
	TransportHandlingUnitTypeCode []CodeType `json:"TransportHandlingUnitTypeCode,omitempty"`
	// A means of transport associated with this transport handling unit.
	TransportMeans []TransportMeansDetails `json:"TransportMeans,omitempty"`
}

// A class for information about pricing structure, lead time, and location associated with
// an item.
type ItemLocationQuantityDetails struct {
	// An allowance or charge associated with this item location quantity.
	AllowanceCharge []AllowanceChargeDetails `json:"AllowanceCharge,omitempty"`
	// A tax category applicable to this item location quantity.
	ApplicableTaxCategory []TaxCategoryDetails `json:"ApplicableTaxCategory,omitempty"`
	// The applicable sales territory.
	ApplicableTerritoryAddress []AddressDetails `json:"ApplicableTerritoryAddress,omitempty"`
	// A delivery unit in which the item is located.
	DeliveryUnit []DeliveryUnitDetails `json:"DeliveryUnit,omitempty"`
	// The price of the item as a percentage of the price of some other item.
	DependentPriceReference []DependentPriceReferenceDetails `json:"DependentPriceReference,omitempty"`
	// An indication that the transported item, as delivered, in the stated quantity to the
	// stated location, is subject to an international regulation concerning the carriage of
	// dangerous goods (true) or not (false).
	HazardousRiskIndicator []IndicatorType `json:"HazardousRiskIndicator,omitempty"`
	// The lead time, i.e., the time taken from the time at which an item is ordered to the time
	// of its delivery.
	LeadTimeMeasure []MeasureType `json:"LeadTimeMeasure,omitempty"`
	// The maximum quantity that can be ordered to qualify for a specific price.
	MaximumQuantity []QuantityType `json:"MaximumQuantity,omitempty"`
	// The minimum quantity that can be ordered to qualify for a specific price.
	MinimumQuantity []QuantityType `json:"MinimumQuantity,omitempty"`
	// The package to which this price applies.
	Package []PackageDetails `json:"Package,omitempty"`
	// The price associated with the given location.
	Price []PriceDetails `json:"Price,omitempty"`
	// Text describing trade restrictions on the quantity of this item or on the item itself.
	TradingRestrictions []TextType `json:"TradingRestrictions,omitempty"`
}

// A reference to the basis for pricing. This may be based on a catalogue or a quoted amount
// from a price list and include some alternative pricing conditions.
type PricingReferenceDetails struct {
	// The price expressed in terms other than the actual price, e.g., the list price v. the
	// contracted price, or the price in bags v. the price in kilos, or the list price in bags
	// v. the contracted price in kilos.
	AlternativeConditionPrice []PriceDetails `json:"AlternativeConditionPrice,omitempty"`
	// An original set of location-specific properties (e.g., price and quantity) associated
	// with this item.
	OriginalItemLocationQuantity []ItemLocationQuantityDetails `json:"OriginalItemLocationQuantity,omitempty"`
}

// A class to describe a price, expressed in a data structure containing multiple properties
// (compare with UnstructuredPrice).
type PriceDetails struct {
	// An allowance or charge associated with this price.
	AllowanceCharge []AllowanceChargeDetails `json:"AllowanceCharge,omitempty"`
	// The quantity at which this price applies.
	BaseQuantity []QuantityType `json:"BaseQuantity,omitempty"`
	// The factor by which the base price unit can be converted to the orderable unit.
	OrderableUnitFactorRate []NumericType `json:"OrderableUnitFactorRate,omitempty"`
	// The amount of the price.
	PriceAmount []AmountType `json:"PriceAmount"`
	// A reason for a price change.
	PriceChangeReason []TextType `json:"PriceChangeReason,omitempty"`
	// Information about a price list applicable to this price.
	PriceList []PriceListDetails `json:"PriceList,omitempty"`
	// The type of price, expressed as text.
	PriceType []TextType `json:"PriceType,omitempty"`
	// The type of price, expressed as a code.
	PriceTypeCode []CodeType `json:"PriceTypeCode,omitempty"`
	// The exchange rate applicable to this price, if it differs from the exchange rate
	// applicable to the document as a whole.
	PricingExchangeRate []ExchangeRateDetails `json:"PricingExchangeRate,omitempty"`
	// A period during which this price is valid.
	ValidityPeriod []PeriodDetails `json:"ValidityPeriod,omitempty"`
}

// A class to define a line in an Invoice.
type InvoiceLineDetails struct {
	// The buyer's accounting cost centre for this invoice line, expressed as text.
	AccountingCost []TextType `json:"AccountingCost,omitempty"`
	// The buyer's accounting cost centre for this invoice line, expressed as a code.
	AccountingCostCode []CodeType `json:"AccountingCostCode,omitempty"`
	// An allowance or charge associated with this invoice line.
	AllowanceCharge []AllowanceChargeDetails `json:"AllowanceCharge,omitempty"`
	// A reference to a billing document associated with this invoice line.
	BillingReference []BillingReferenceDetails `json:"BillingReference,omitempty"`
	// A delivery associated with this invoice line.
	Delivery []DeliveryDetails `json:"Delivery,omitempty"`
	// Terms and conditions of the delivery associated with this invoice line.
	DeliveryTerms []DeliveryTermsDetails `json:"DeliveryTerms,omitempty"`
	// A reference to a despatch line associated with this invoice line.
	DespatchLineReference []LineReferenceDetails `json:"DespatchLineReference,omitempty"`
	// A reference to a document associated with this invoice line.
	DocumentReference []DocumentReferenceDetails `json:"DocumentReference,omitempty"`
	// An indicator that this invoice line is free of charge (true) or not (false). The default
	// is false.
	FreeOfChargeIndicator []IndicatorType `json:"FreeOfChargeIndicator,omitempty"`
	// An identifier for this invoice line.
	ID []IdentifierType `json:"ID"`
	// The quantity (of items) on this invoice line.
	InvoicedQuantity []QuantityType `json:"InvoicedQuantity,omitempty"`
	// An invoice period to which this invoice line applies.
	InvoicePeriod []PeriodDetails `json:"InvoicePeriod,omitempty"`
	// The item associated with this invoice line.
	Item []ItemDetails `json:"Item"`
	// The price extension, calculated by multiplying the price per unit by the quantity of
	// items on this invoice line.
	ItemPriceExtension []PriceExtensionDetails `json:"ItemPriceExtension,omitempty"`
	// The total amount for this invoice line, including allowance charges but net of taxes.
	LineExtensionAmount []AmountType `json:"LineExtensionAmount"`
	// Free-form text conveying information that is not contained explicitly in other structures.
	Note []TextType `json:"Note,omitempty"`
	// A reference to an order line associated with this invoice line.
	OrderLineReference []OrderLineReferenceDetails `json:"OrderLineReference,omitempty"`
	// The party who originated the Order to which the Invoice is related.
	OriginatorParty []PartyDetails `json:"OriginatorParty,omitempty"`
	// A code signifying the business purpose for this payment.
	PaymentPurposeCode []CodeType `json:"PaymentPurposeCode,omitempty"`
	// A specification of payment terms associated with this invoice line.
	PaymentTerms []PaymentTermsDetails `json:"PaymentTerms,omitempty"`
	// The price of the item associated with this invoice line.
	Price []PriceDetails `json:"Price,omitempty"`
	// A reference to pricing and item location information associated with this invoice line.
	PricingReference []PricingReferenceDetails `json:"PricingReference,omitempty"`
	// A reference to a receipt line associated with this invoice line.
	ReceiptLineReference []LineReferenceDetails `json:"ReceiptLineReference,omitempty"`
	// An invoice line subsidiary to this invoice line.
	SubInvoiceLine []InvoiceLineDetails `json:"SubInvoiceLine,omitempty"`
	// The date of this invoice line, used to indicate the point at which tax becomes applicable.
	TaxPointDate []DateType `json:"TaxPointDate,omitempty"`
	// A total amount of taxes of a particular kind applicable to this invoice line.
	TaxTotal []TaxTotalDetails `json:"TaxTotal,omitempty"`
	// A universally unique identifier for this invoice line.
	UUID []IdentifierType `json:"UUID,omitempty"`
	// A reference to a TaxTotal class describing the amount that has been withhold by the
	// authorities, e.g. if the creditor is in dept because of non paid taxes.
	WithholdingTaxTotal []TaxTotalDetails `json:"WithholdingTaxTotal,omitempty"`
}

// A class defining how goods items are split across transport equipment.
type GoodsItemContainerDetails struct {
	// An identifier for this goods item container.
	ID []IdentifierType `json:"ID"`
	// The number of goods items loaded into or onto one piece of transport equipment as a total
	// consignment or part of a consignment.
	Quantity []QuantityType `json:"Quantity,omitempty"`
	// A piece of transport equipment used to contain a single goods item.
	TransportEquipment []TransportEquipmentDetails `json:"TransportEquipment,omitempty"`
}

// A class to describe a package.
type PackageDetails struct {
	// A package contained within this package.
	ContainedPackage []PackageDetails `json:"ContainedPackage,omitempty"`
	// The piece of transport equipment containing this package.
	ContainingTransportEquipment []TransportEquipmentDetails `json:"ContainingTransportEquipment,omitempty"`
	// The delivery of this package.
	Delivery []DeliveryDetails `json:"Delivery,omitempty"`
	// A delivery unit within this package.
	DeliveryUnit []DeliveryUnitDetails `json:"DeliveryUnit,omitempty"`
	// The despatch of this package.
	Despatch []DespatchDetails `json:"Despatch,omitempty"`
	// A goods item included in this package.
	GoodsItem []GoodsItemDetails `json:"GoodsItem,omitempty"`
	// An identifier for this package.
	ID []IdentifierType `json:"ID,omitempty"`
	// A measurable dimension (length, mass, weight, or volume) of this package.
	MeasurementDimension []DimensionDetails `json:"MeasurementDimension,omitempty"`
	// A code signifying a level of packaging.
	PackageLevelCode []CodeType `json:"PackageLevelCode,omitempty"`
	// A code signifying a type of packaging.
	PackagingTypeCode []CodeType `json:"PackagingTypeCode,omitempty"`
	// Text describing the packaging material.
	PackingMaterial []TextType `json:"PackingMaterial,omitempty"`
	// The pickup of this package.
	Pickup []PickupDetails `json:"Pickup,omitempty"`
	// The quantity of items contained in this package.
	Quantity []QuantityType `json:"Quantity,omitempty"`
	// An indicator that the packaging material is returnable (true) or not (false).
	ReturnableMaterialIndicator []IndicatorType `json:"ReturnableMaterialIndicator,omitempty"`
	// An identifier for use in tracing this package, such as the EPC number used in RFID.
	TraceID []IdentifierType `json:"TraceID,omitempty"`
}

// A class to describe a separately identifiable quantity of goods of a single product type.
type GoodsItemDetails struct {
	// The number of units in the goods item to which charges apply.
	ChargeableQuantity []QuantityType `json:"ChargeableQuantity,omitempty"`
	// The weight on which a charge is to be based.
	ChargeableWeightMeasure []MeasureType `json:"ChargeableWeightMeasure,omitempty"`
	// A goods item contained in this goods item.
	ContainedGoodsItem []GoodsItemDetails `json:"ContainedGoodsItem,omitempty"`
	// A package containing this goods item.
	ContainingPackage []PackageDetails `json:"ContainingPackage,omitempty"`
	// An indicator that this goods item has been classified for import by customs (true) or not
	// (false).
	CustomsImportClassifiedIndicator []IndicatorType `json:"CustomsImportClassifiedIndicator,omitempty"`
	// A code assigned by customs to signify the status of this goods item.
	CustomsStatusCode []CodeType `json:"CustomsStatusCode,omitempty"`
	// Quantity of the units in this goods item as required by customs for tariff, statistical,
	// or fiscal purposes.
	CustomsTariffQuantity []QuantityType `json:"CustomsTariffQuantity,omitempty"`
	// The total declared value for customs purposes of the goods item.
	DeclaredCustomsValueAmount []AmountType `json:"DeclaredCustomsValueAmount,omitempty"`
	// The value of this goods item, declared by the shipper or his agent solely for the purpose
	// of varying the carrier's level of liability from that provided in the contract of
	// carriage, in case of loss or damage to goods or delayed delivery.
	DeclaredForCarriageValueAmount []AmountType `json:"DeclaredForCarriageValueAmount,omitempty"`
	// The total declared value of all the goods items in the same consignment with this goods
	// item that have the same statistical heading.
	DeclaredStatisticsValueAmount []AmountType `json:"DeclaredStatisticsValueAmount,omitempty"`
	// The delivery of this goods item.
	Delivery []DeliveryDetails `json:"Delivery,omitempty"`
	// Text describing this goods item to identify it for customs, statistical, or transport
	// purposes.
	Description []TextType `json:"Description,omitempty"`
	// The despatch of this goods item.
	Despatch []DespatchDetails `json:"Despatch,omitempty"`
	// The monetary amount that has to be or has been paid as calculated under the applicable
	// trade delivery.
	FreeOnBoardValueAmount []AmountType `json:"FreeOnBoardValueAmount,omitempty"`
	// A cost incurred by the shipper in moving goods, by whatever means, from one place to
	// another under the terms of the contract of carriage. In addition to transport costs, this
	// may include such elements as packing, documentation, loading, unloading, and insurance to
	// the extent that they relate to the freight costs.
	FreightAllowanceCharge []AllowanceChargeDetails `json:"FreightAllowanceCharge,omitempty"`
	// The transporting of a goods item in a unit of transport equipment (e.g., container).
	GoodsItemContainer []GoodsItemContainerDetails `json:"GoodsItemContainer,omitempty"`
	// The volume of this goods item, normally calculated by multiplying its maximum length,
	// width, and height.
	GrossVolumeMeasure []MeasureType `json:"GrossVolumeMeasure,omitempty"`
	// The weight of this goods item, including packing and packaging but excluding the
	// carrier's equipment.
	GrossWeightMeasure []MeasureType `json:"GrossWeightMeasure,omitempty"`
	// An indication that the transported goods item is subject to an international regulation
	// concerning the carriage of dangerous goods (true) or not (false).
	HazardousRiskIndicator []IndicatorType `json:"HazardousRiskIndicator,omitempty"`
	// An identifier for this goods item.
	ID []IdentifierType `json:"ID,omitempty"`
	// The amount covered by insurance for this goods item.
	InsuranceValueAmount []AmountType `json:"InsuranceValueAmount,omitempty"`
	// Information about an invoice line relating to this goods item.
	InvoiceLine []InvoiceLineDetails `json:"InvoiceLine,omitempty"`
	// Product information relating to a goods item.
	Item []ItemDetails `json:"Item,omitempty"`
	// Information about maximum temperature.
	MaximumTemperature []TemperatureDetails `json:"MaximumTemperature,omitempty"`
	// A measurable dimension (length, mass, weight, or volume) of this goods item.
	MeasurementDimension []DimensionDetails `json:"MeasurementDimension,omitempty"`
	// Information about minimum temperature.
	MinimumTemperature []TemperatureDetails `json:"MinimumTemperature,omitempty"`
	// The total weight of this goods item, excluding all packing and packaging.
	NetNetWeightMeasure []MeasureType `json:"NetNetWeightMeasure,omitempty"`
	// The volume contained by a goods item, excluding the volume of any packaging material.
	NetVolumeMeasure []MeasureType `json:"NetVolumeMeasure,omitempty"`
	// The weight of this goods item, excluding packing but including packaging that normally
	// accompanies the goods.
	NetWeightMeasure []MeasureType `json:"NetWeightMeasure,omitempty"`
	// The region in which the goods have been produced or manufactured, according to criteria
	// laid down for the purposes of application of the customs tariff, or of quantitative
	// restrictions, or of any other measure related to trade.
	OriginAddress []AddressDetails `json:"OriginAddress,omitempty"`
	// The pickup of this goods item.
	Pickup []PickupDetails `json:"Pickup,omitempty"`
	// A code signifying the treatment preference for this goods item according to international
	// trading agreements.
	PreferenceCriterionCode []CodeType `json:"PreferenceCriterionCode,omitempty"`
	// The number of units making up this goods item.
	Quantity []QuantityType `json:"Quantity,omitempty"`
	// An identifier for a set of tariff codes required to specify a type of goods for customs,
	// transport, statistical, or other regulatory purposes.
	RequiredCustomsID []IdentifierType `json:"RequiredCustomsID,omitempty"`
	// The number of units in the goods item that may be returned.
	ReturnableQuantity []QuantityType `json:"ReturnableQuantity,omitempty"`
	// A sequence number differentiating a specific goods item within a consignment.
	SequenceNumberID []IdentifierType `json:"SequenceNumberID,omitempty"`
	// A reference to a shipping document associated with this goods item.
	ShipmentDocumentReference []DocumentReferenceDetails `json:"ShipmentDocumentReference,omitempty"`
	// The temperature of the goods item.
	Temperature []TemperatureDetails `json:"Temperature,omitempty"`
	// An identifier for use in tracing this goods item, such as the EPC number used in RFID.
	TraceID []IdentifierType `json:"TraceID,omitempty"`
	// The amount on which a duty, tax, or fee will be assessed.
	ValueAmount []AmountType `json:"ValueAmount,omitempty"`
}

// A class to describe a piece of equipment used to transport goods.
type TransportEquipmentDetails struct {
	// The percent of the airflow within this piece of transport equipment.
	AirFlowPercent []NumericType `json:"AirFlowPercent,omitempty"`
	// An indicator that this piece of transport equipment is approved for animal food (true) or
	// not (false).
	AnimalFoodApprovedIndicator []IndicatorType `json:"AnimalFoodApprovedIndicator,omitempty"`
	// The applicable transport means associated with this piece of transport equipment.
	ApplicableTransportMeans []TransportMeansDetails `json:"ApplicableTransportMeans,omitempty"`
	// A piece of transport equipment attached to this piece of transport equipment.
	AttachedTransportEquipment []TransportEquipmentDetails `json:"AttachedTransportEquipment,omitempty"`
	// Characteristics of this piece of transport equipment.
	Characteristics []TextType `json:"Characteristics,omitempty"`
	// A piece of transport equipment contained in this piece of transport equipment.
	ContainedInTransportEquipment []TransportEquipmentDetails `json:"ContainedInTransportEquipment,omitempty"`
	// Damage associated with this piece of transport equipment.
	DamageRemarks []TextType `json:"DamageRemarks,omitempty"`
	// An indicator that this piece of transport equipment is approved for dangerous goods
	// (true) or not (false).
	DangerousGoodsApprovedIndicator []IndicatorType `json:"DangerousGoodsApprovedIndicator,omitempty"`
	// The delivery of this piece of transport equipment.
	Delivery []DeliveryDetails `json:"Delivery,omitempty"`
	// A delivery of this piece of transport equipment.
	DeliveryTransportEvent []TransportEventDetails `json:"DeliveryTransportEvent,omitempty"`
	// Text describing this piece of transport equipment.
	Description []TextType `json:"Description,omitempty"`
	// The despatch of this piece of transport equipment.
	Despatch []DespatchDetails `json:"Despatch,omitempty"`
	// A code signifying the current disposition of this piece of transport equipment.
	DispositionCode []CodeType `json:"DispositionCode,omitempty"`
	// A freight allowance charge associated with this piece of transport equipment.
	FreightAllowanceCharge []AllowanceChargeDetails `json:"FreightAllowanceCharge,omitempty"`
	// A code signifying whether this piece of transport equipment is full, partially full, or
	// empty.
	FullnessIndicationCode []CodeType `json:"FullnessIndicationCode,omitempty"`
	// A goods item contained in this piece of transport equipment.
	GoodsItem []GoodsItemDetails `json:"GoodsItem,omitempty"`
	// The gross volume of this piece of transport equipment.
	GrossVolumeMeasure []MeasureType `json:"GrossVolumeMeasure,omitempty"`
	// The gross weight of this piece of transport equipment.
	GrossWeightMeasure []MeasureType `json:"GrossWeightMeasure,omitempty"`
	// A handling of this piece of transport equipment.
	HandlingTransportEvent []TransportEventDetails `json:"HandlingTransportEvent,omitempty"`
	// A set of haulage trading terms associated with this piece of transport equipment.
	HaulageTradingTerms []TradingTermsDetails `json:"HaulageTradingTerms,omitempty"`
	// Transit-related information regarding a type of hazardous goods contained in this piece
	// of transport equipment.
	HazardousGoodsTransit []HazardousGoodsTransitDetails `json:"HazardousGoodsTransit,omitempty"`
	// An indicator that this piece of transport equipment is approved for human food (true) or
	// not (false).
	HumanFoodApprovedIndicator []IndicatorType `json:"HumanFoodApprovedIndicator,omitempty"`
	// The percent humidity within this piece of transport equipment.
	HumidityPercent []NumericType `json:"HumidityPercent,omitempty"`
	// An identifier for this piece of transport equipment.
	ID []IdentifierType `json:"ID,omitempty"`
	// Additional information about this piece of transport equipment.
	Information []TextType `json:"Information,omitempty"`
	// An indication of the legal status of this piece of transport equipment with respect to
	// the Container Convention Code.
	LegalStatusIndicator []IndicatorType `json:"LegalStatusIndicator,omitempty"`
	// The location where this piece of transport equipment is loaded.
	LoadingLocation []LocationDetails `json:"LoadingLocation,omitempty"`
	// The authorized party responsible for certifying that the goods were loaded into this
	// piece of transport equipment.
	LoadingProofParty []PartyDetails `json:"LoadingProofParty,omitempty"`
	// A loading of this piece of transport equipment.
	LoadingTransportEvent []TransportEventDetails `json:"LoadingTransportEvent,omitempty"`
	// In the case of a refrigeration unit, the maximum allowable operating temperature for this
	// container.
	MaximumTemperature []TemperatureDetails `json:"MaximumTemperature,omitempty"`
	// A measurable dimension (length, mass, weight, or volume) of this piece of transport
	// equipment.
	MeasurementDimension []DimensionDetails `json:"MeasurementDimension,omitempty"`
	// In the case of a refrigeration unit, the minimum allowable operating temperature for this
	// container.
	MinimumTemperature []TemperatureDetails `json:"MinimumTemperature,omitempty"`
	// The party that operates this piece of transport equipment.
	OperatingParty []PartyDetails `json:"OperatingParty,omitempty"`
	// The party that owns this piece of transport equipment.
	OwnerParty []PartyDetails `json:"OwnerParty,omitempty"`
	// A code signifying the type of owner of this piece of transport equipment.
	OwnerTypeCode []CodeType `json:"OwnerTypeCode,omitempty"`
	// A package contained in this piece of transport equipment.
	Package []PackageDetails `json:"Package,omitempty"`
	// A packaged transport handling unit associated with this piece of transport equipment.
	PackagedTransportHandlingUnit []TransportHandlingUnitDetails `json:"PackagedTransportHandlingUnit,omitempty"`
	// The pickup of this piece of transport equipment.
	Pickup []PickupDetails `json:"Pickup,omitempty"`
	// A pickup of this piece of transport equipment.
	PickupTransportEvent []TransportEventDetails `json:"PickupTransportEvent,omitempty"`
	// A positioning of this piece of transport equipment.
	PositioningTransportEvent []TransportEventDetails `json:"PositioningTransportEvent,omitempty"`
	// An indicator that this piece of transport equipment can supply power (true) or not
	// (false).
	PowerIndicator []IndicatorType `json:"PowerIndicator,omitempty"`
	// The party providing this piece of transport equipment.
	ProviderParty []PartyDetails `json:"ProviderParty,omitempty"`
	// A code signifying the type of provider of this piece of transport equipment.
	ProviderTypeCode []CodeType `json:"ProviderTypeCode,omitempty"`
	// A quarantine of this piece of transport equipment.
	QuarantineTransportEvent []TransportEventDetails `json:"QuarantineTransportEvent,omitempty"`
	// An identifier for the consignment contained by this piece of transport equipment.
	ReferencedConsignmentID []IdentifierType `json:"ReferencedConsignmentID,omitempty"`
	// An indicator that this piece of transport equipment is refrigerated (true) or not (false).
	RefrigeratedIndicator []IndicatorType `json:"RefrigeratedIndicator,omitempty"`
	// An indicator that this piece of transport equipment's refrigeration is on (true) or off
	// (false).
	RefrigerationOnIndicator []IndicatorType `json:"RefrigerationOnIndicator,omitempty"`
	// An indicator that this piece of transport equipment is returnable (true) or not (false).
	ReturnabilityIndicator []IndicatorType `json:"ReturnabilityIndicator,omitempty"`
	// A service allowance charge associated with this piece of transport equipment.
	ServiceAllowanceCharge []AllowanceChargeDetails `json:"ServiceAllowanceCharge,omitempty"`
	// A reference to a shipping document associated with this piece of transport equipment.
	ShipmentDocumentReference []DocumentReferenceDetails `json:"ShipmentDocumentReference,omitempty"`
	// A code signifying the size and type of this piece of piece of transport equipment. When
	// the piece of transport equipment is a shipping container, it is recommended to use
	// ContainerSizeTypeCode for validation.
	SizeTypeCode []CodeType `json:"SizeTypeCode,omitempty"`
	// Special transport requirements expressed as text.
	SpecialTransportRequirements []TextType `json:"SpecialTransportRequirements,omitempty"`
	// The location where this piece of transport equipment is being stored.
	StorageLocation []LocationDetails `json:"StorageLocation,omitempty"`
	// The party that supplies this piece of transport equipment.
	SupplierParty []SupplierPartyDetails `json:"SupplierParty,omitempty"`
	// The weight of this piece of transport equipment when empty.
	TareWeightMeasure []MeasureType `json:"TareWeightMeasure,omitempty"`
	// An identifier for use in tracing this piece of transport equipment, such as the EPC
	// number used in RFID.
	TraceID []IdentifierType `json:"TraceID,omitempty"`
	// A code signifying the tracking device for this piece of transport equipment.
	TrackingDeviceCode []CodeType `json:"TrackingDeviceCode,omitempty"`
	// A seal securing the door of a piece of transport equipment.
	TransportEquipmentSeal []TransportEquipmentSealDetails `json:"TransportEquipmentSeal,omitempty"`
	// A code signifying the type of this piece of transport equipment.
	TransportEquipmentTypeCode []CodeType `json:"TransportEquipmentTypeCode,omitempty"`
	// A transport event associated with this piece of transport equipment.
	TransportEvent []TransportEventDetails `json:"TransportEvent,omitempty"`
	// The location where this piece of transport equipment is unloaded.
	UnloadingLocation []LocationDetails `json:"UnloadingLocation,omitempty"`
}

// A class to describe a significant occurrence or happening related to the transportation
// of goods.
type TransportEventDetails struct {
	// An indicator that this transport event has been completed (true) or not (false).
	CompletionIndicator []IndicatorType `json:"CompletionIndicator,omitempty"`
	// A contact associated with this transport event.
	Contact []ContactDetails `json:"Contact,omitempty"`
	// The current status of this transport event.
	CurrentStatus []StatusDetails `json:"CurrentStatus,omitempty"`
	// Text describing this transport event.
	Description []TextType `json:"Description,omitempty"`
	// An identifier for this transport event within an agreed event identification scheme.
	IdentificationID []IdentifierType `json:"IdentificationID,omitempty"`
	// The location associated with this transport event.
	Location []LocationDetails `json:"Location,omitempty"`
	// The date of this transport event.
	OccurrenceDate []DateType `json:"OccurrenceDate,omitempty"`
	// The time of this transport event.
	OccurrenceTime []TimeType `json:"OccurrenceTime,omitempty"`
	// A period of time associated with this transport event.
	Period []PeriodDetails `json:"Period,omitempty"`
	// The shipment involved in this transport event.
	ReportedShipment []ShipmentDetails `json:"ReportedShipment,omitempty"`
	// A signature that can be used to sign for an entry or an exit at a transport location
	// (e.g., port terminal).
	Signature []SignatureDetails `json:"Signature,omitempty"`
	// A code signifying the type of this transport event.
	TransportEventTypeCode []CodeType `json:"TransportEventTypeCode,omitempty"`
}

// A class to describe one stage of movement in a transport of goods.
type ShipmentStageDetails struct {
	// The acceptance of goods in this shipment stage.
	AcceptanceTransportEvent []TransportEventDetails `json:"AcceptanceTransportEvent,omitempty"`
	// The actual arrival at a specific location during a transportation service.
	ActualArrivalTransportEvent []TransportEventDetails `json:"ActualArrivalTransportEvent,omitempty"`
	// The actual departure from a specific location during a transportation service.
	ActualDepartureTransportEvent []TransportEventDetails `json:"ActualDepartureTransportEvent,omitempty"`
	// The pickup of goods in this shipment stage.
	ActualPickupTransportEvent []TransportEventDetails `json:"ActualPickupTransportEvent,omitempty"`
	// The location of an actual waypoint during a transportation service.
	ActualWaypointTransportEvent []TransportEventDetails `json:"ActualWaypointTransportEvent,omitempty"`
	// The making available of shipments in this shipment stage.
	AvailabilityTransportEvent []TransportEventDetails `json:"AvailabilityTransportEvent,omitempty"`
	// A carrier party responsible for this shipment stage.
	CarrierParty []PartyDetails `json:"CarrierParty,omitempty"`
	// A person operating or serving aboard a transport means.
	CrewMemberPerson []PersonDetails `json:"CrewMemberPerson,omitempty"`
	// The total number of crew aboard a transport means.
	CrewQuantity []QuantityType `json:"CrewQuantity,omitempty"`
	// A customs agent associated with this shipment stage.
	CustomsAgentParty []PartyDetails `json:"CustomsAgentParty,omitempty"`
	// The delivery of goods in this shipment stage.
	DeliveryTransportEvent []TransportEventDetails `json:"DeliveryTransportEvent,omitempty"`
	// Text of instructions relating to demurrage (the case in which a vessel is prevented from
	// loading or discharging cargo within the stipulated laytime).
	DemurrageInstructions []TextType `json:"DemurrageInstructions,omitempty"`
	// The detention of a transport means during loading and unloading operations.
	DetentionTransportEvent []TransportEventDetails `json:"DetentionTransportEvent,omitempty"`
	// The discharge event associated with this shipment stage.
	DischargeTransportEvent []TransportEventDetails `json:"DischargeTransportEvent,omitempty"`
	// Describes a person responsible for driving the transport means.
	DriverPerson []PersonDetails `json:"DriverPerson,omitempty"`
	// The dropping off of goods in this shipment stage.
	DropoffTransportEvent []TransportEventDetails `json:"DropoffTransportEvent,omitempty"`
	// Describes an estimated arrival at a location during a transport service.
	EstimatedArrivalTransportEvent []TransportEventDetails `json:"EstimatedArrivalTransportEvent,omitempty"`
	// The estimated date of delivery in this shipment stage.
	EstimatedDeliveryDate []DateType `json:"EstimatedDeliveryDate,omitempty"`
	// The estimated time of delivery in this shipment stage.
	EstimatedDeliveryTime []TimeType `json:"EstimatedDeliveryTime,omitempty"`
	// Describes an estimated departure at a location during a transport service.
	EstimatedDepartureTransportEvent []TransportEventDetails `json:"EstimatedDepartureTransportEvent,omitempty"`
	// The estimated transit period of this shipment stage.
	EstimatedTransitPeriod []PeriodDetails `json:"EstimatedTransitPeriod,omitempty"`
	// The examination of shipments in this shipment stage.
	ExaminationTransportEvent []TransportEventDetails `json:"ExaminationTransportEvent,omitempty"`
	// The export event associated with this shipment stage.
	ExportationTransportEvent []TransportEventDetails `json:"ExportationTransportEvent,omitempty"`
	// A freight allowance charge for this shipment stage.
	FreightAllowanceCharge []AllowanceChargeDetails `json:"FreightAllowanceCharge,omitempty"`
	// The location associated with a freight charge related to this shipment stage.
	FreightChargeLocation []LocationDetails `json:"FreightChargeLocation,omitempty"`
	// An identifier for this shipment stage.
	ID []IdentifierType `json:"ID,omitempty"`
	// Text of instructions applicable to a shipment stage.
	Instructions []TextType `json:"Instructions,omitempty"`
	// The location of loading for a shipment stage.
	LoadingPortLocation []LocationDetails `json:"LoadingPortLocation,omitempty"`
	// An identifier for the loading sequence (of consignments) associated with this shipment
	// stage.
	LoadingSequenceID []IdentifierType `json:"LoadingSequenceID,omitempty"`
	// The loading of goods in this shipment stage.
	LoadingTransportEvent []TransportEventDetails `json:"LoadingTransportEvent,omitempty"`
	// The person responsible for the ship's safe and efficient operation, including cargo
	// operations, navigation, crew management and for ensuring that the vessel complies with
	// local and international laws, as well as company and flag state policies.
	MasterPerson []PersonDetails `json:"MasterPerson,omitempty"`
	// An indicator that this stage takes place after the main carriage of the shipment (true)
	// or not (false).
	OnCarriageIndicator []IndicatorType `json:"OnCarriageIndicator,omitempty"`
	// The optional takeover of the goods in this shipment stage.
	OptionalTakeoverTransportEvent []TransportEventDetails `json:"OptionalTakeoverTransportEvent,omitempty"`
	// A person who travels in a conveyance without participating in its operation.
	PassengerPerson []PersonDetails `json:"PassengerPerson,omitempty"`
	// The total number of passengers aboard a transport means.
	PassengerQuantity []QuantityType `json:"PassengerQuantity,omitempty"`
	// The arrival planned by the party providing a transportation service.
	PlannedArrivalTransportEvent []TransportEventDetails `json:"PlannedArrivalTransportEvent,omitempty"`
	// The departure planned by the party providing a transportation service.
	PlannedDepartureTransportEvent []TransportEventDetails `json:"PlannedDepartureTransportEvent,omitempty"`
	// A waypoint planned by the party providing a transportation service.
	PlannedWaypointTransportEvent []TransportEventDetails `json:"PlannedWaypointTransportEvent,omitempty"`
	// An indicator that this stage takes place before the main carriage of the shipment (true)
	// or not (false).
	PreCarriageIndicator []IndicatorType `json:"PreCarriageIndicator,omitempty"`
	// The receipt of goods in this shipment stage.
	ReceiptTransportEvent []TransportEventDetails `json:"ReceiptTransportEvent,omitempty"`
	// Describes a person being responsible for providing the required administrative reporting
	// relating to a transport.
	ReportingPerson []PersonDetails `json:"ReportingPerson,omitempty"`
	// The arrival requested by the party requesting a transportation service.
	RequestedArrivalTransportEvent []TransportEventDetails `json:"RequestedArrivalTransportEvent,omitempty"`
	// The departure requested by the party requesting a transportation service.
	RequestedDepartureTransportEvent []TransportEventDetails `json:"RequestedDepartureTransportEvent,omitempty"`
	// A waypoint requested by the party requesting a transportation service.
	RequestedWaypointTransportEvent []TransportEventDetails `json:"RequestedWaypointTransportEvent,omitempty"`
	// The delivery date required by the buyer in this shipment stage.
	RequiredDeliveryDate []DateType `json:"RequiredDeliveryDate,omitempty"`
	// The delivery time required by the buyer in this shipment stage.
	RequiredDeliveryTime []TimeType `json:"RequiredDeliveryTime,omitempty"`
	// The person on board the vessel, accountable to the master, designated by the company as
	// responsible for the security of the ship, including implementation and maintenance of the
	// ship security plan and for the liaison with the company security officer and the port
	// facility security officers.
	SecurityOfficerPerson []PersonDetails `json:"SecurityOfficerPerson,omitempty"`
	// The person responsible for the health of the people aboard a ship at sea.
	ShipsSurgeonPerson []PersonDetails `json:"ShipsSurgeonPerson,omitempty"`
	// The storage of goods in this shipment stage.
	StorageTransportEvent []TransportEventDetails `json:"StorageTransportEvent,omitempty"`
	// Identifies the successive loading sequence (of consignments) associated with a shipment
	// stage.
	SuccessiveSequenceID []IdentifierType `json:"SuccessiveSequenceID,omitempty"`
	// The receiver's takeover of the goods in this shipment stage.
	TakeoverTransportEvent []TransportEventDetails `json:"TakeoverTransportEvent,omitempty"`
	// A terminal operator associated with this shipment stage.
	TerminalOperatorParty []PartyDetails `json:"TerminalOperatorParty,omitempty"`
	// A code signifying the direction of transit in this shipment stage.
	TransitDirectionCode []CodeType `json:"TransitDirectionCode,omitempty"`
	// The period during which this shipment stage actually took place.
	TransitPeriod []PeriodDetails `json:"TransitPeriod,omitempty"`
	// A significant occurrence in the course of this shipment of goods.
	TransportEvent []TransportEventDetails `json:"TransportEvent,omitempty"`
	// The means of transport used in this shipment stage.
	TransportMeans []TransportMeansDetails `json:"TransportMeans,omitempty"`
	// A code signifying the kind of transport means (truck, vessel, etc.) used for this
	// shipment stage.
	TransportMeansTypeCode []CodeType `json:"TransportMeansTypeCode,omitempty"`
	// A code signifying the method of transport used for this shipment stage.
	TransportModeCode []CodeType `json:"TransportModeCode,omitempty"`
	// The location of transshipment relating to a shipment stage.
	TransshipPortLocation []LocationDetails `json:"TransshipPortLocation,omitempty"`
	// The location of unloading for a shipment stage.
	UnloadingPortLocation []LocationDetails `json:"UnloadingPortLocation,omitempty"`
	// The warehousing event associated with this shipment stage.
	WarehousingTransportEvent []TransportEventDetails `json:"WarehousingTransportEvent,omitempty"`
}

// A class to describe a transportation service.
type TransportationServiceDetails struct {
	// A classification of this transportation service.
	CommodityClassification []CommodityClassificationDetails `json:"CommodityClassification,omitempty"`
	// An environmental emission resulting from this transportation service.
	EnvironmentalEmission []EnvironmentalEmissionDetails `json:"EnvironmentalEmission,omitempty"`
	// The estimated duration of this transportation service.
	EstimatedDurationPeriod []PeriodDetails `json:"EstimatedDurationPeriod,omitempty"`
	// A code signifying the rate class for freight in this transportation service.
	FreightRateClassCode []CodeType `json:"FreightRateClassCode,omitempty"`
	// The name of this transportation service.
	Name []TextType `json:"Name,omitempty"`
	// In a transport contract, the deadline date by which this transportation service has to be
	// booked. For example, if this service is scheduled for Wednesday 16 February 2011 at 10
	// a.m. CET, the nomination date might be Tuesday15 February 2011.
	NominationDate []DateType `json:"NominationDate,omitempty"`
	// In a transport contract, the deadline time by which this transportation service has to be
	// booked. For example, if this service is scheduled for Wednesday 16 February 2011 at 10
	// a.m. CET, the nomination date might be Tuesday15 February 2011 and the nomination time 4
	// p.m. at the latest.
	NominationTime []TimeType `json:"NominationTime,omitempty"`
	// The priority of this transportation service.
	Priority []TextType `json:"Priority,omitempty"`
	// The transport service provider responsible for this transportation service.
	ResponsibleTransportServiceProviderParty []PartyDetails `json:"ResponsibleTransportServiceProviderParty,omitempty"`
	// A class to specify which day of the week a transport service is operational.
	ScheduledServiceFrequency []ServiceFrequencyDetails `json:"ScheduledServiceFrequency,omitempty"`
	// A number indicating the order of this transportation service in a sequence of
	// transportation services.
	SequenceNumeric []NumericType `json:"SequenceNumeric,omitempty"`
	// One of the stages of shipment in this transportation service.
	ShipmentStage []ShipmentStageDetails `json:"ShipmentStage,omitempty"`
	// A classification (e.g., general cargo) for commodities that can be handled in this
	// transportation service.
	SupportedCommodityClassification []CommodityClassificationDetails `json:"SupportedCommodityClassification,omitempty"`
	// A piece of transport equipment supported in this transportation service.
	SupportedTransportEquipment []TransportEquipmentDetails `json:"SupportedTransportEquipment,omitempty"`
	// A code signifying the tariff class applicable to this transportation service.
	TariffClassCode []CodeType `json:"TariffClassCode,omitempty"`
	// The total capacity or volume available in this transportation service.
	TotalCapacityDimension []DimensionDetails `json:"TotalCapacityDimension,omitempty"`
	// Text describing this transportation service.
	TransportationServiceDescription []TextType `json:"TransportationServiceDescription,omitempty"`
	// The Uniform Resource Identifier (URI) of a document providing additional details
	// regarding this transportation service.
	TransportationServiceDetailsURI []IdentifierType `json:"TransportationServiceDetailsURI,omitempty"`
	// A piece of transport equipment used in this transportation service.
	TransportEquipment []TransportEquipmentDetails `json:"TransportEquipment,omitempty"`
	// One of the transport events taking place in this transportation service.
	TransportEvent []TransportEventDetails `json:"TransportEvent,omitempty"`
	// A code signifying the extent of this transportation service (e.g., door-to-door,
	// port-to-port).
	TransportServiceCode []CodeType `json:"TransportServiceCode"`
	// A classification for commodities that cannot be handled in this transportation service.
	UnsupportedCommodityClassification []CommodityClassificationDetails `json:"UnsupportedCommodityClassification,omitempty"`
	// A piece of transport equipment that is not supported in this transportation service.
	UnsupportedTransportEquipment []TransportEquipmentDetails `json:"UnsupportedTransportEquipment,omitempty"`
}

// A class to describe a contract.
type ContractDetails struct {
	// A reference to a contract document.
	ContractDocumentReference []DocumentReferenceDetails `json:"ContractDocumentReference,omitempty"`
	// The type of this contract, expressed as text, such as "Cost plus award fee" and "Cost
	// plus fixed fee" from UNCEFACT Contract Type code list.
	ContractType []TextType `json:"ContractType,omitempty"`
	// The type of this contract, expressed as a code, such as "Cost plus award fee" and "Cost
	// plus fixed fee" from UNCEFACT Contract Type code list.
	ContractTypeCode []CodeType `json:"ContractTypeCode,omitempty"`
	// In a transportation contract, the delivery of the services required to book the services
	// specified in the contract.
	ContractualDelivery []DeliveryDetails `json:"ContractualDelivery,omitempty"`
	// Text describing this contract.
	Description []TextType `json:"Description,omitempty"`
	// An identifier for this contract.
	ID []IdentifierType `json:"ID,omitempty"`
	// The date on which this contract was issued.
	IssueDate []DateType `json:"IssueDate,omitempty"`
	// The time at which this contract was issued.
	IssueTime []TimeType `json:"IssueTime,omitempty"`
	// In a transportation contract, the deadline date by which the services referred to in the
	// transport execution plan have to be booked. For example, if this service is a carrier
	// service scheduled for Wednesday 16 February 2011 at 10 a.m. CET, the nomination date
	// might be Tuesday15 February 2011.
	NominationDate []DateType `json:"NominationDate,omitempty"`
	// In a transportation contract, the period required to book the services specified in the
	// contract before the services can begin.
	NominationPeriod []PeriodDetails `json:"NominationPeriod,omitempty"`
	// In a transportation contract, the deadline time by which the services referred to in the
	// transport execution plan have to be booked. For example, if this service is a carrier
	// service scheduled for Wednesday 16 February 2011 at 10 a.m. CET, the nomination date
	// might be Tuesday15 February 2011 and the nomination time 4 p.m. at the latest.
	NominationTime []TimeType `json:"NominationTime,omitempty"`
	// Free-form text conveying information that is not contained explicitly in other structures.
	Note []TextType `json:"Note,omitempty"`
	// The period during which this contract is valid.
	ValidityPeriod []PeriodDetails `json:"ValidityPeriod,omitempty"`
	// An identifier for the current version of this contract.
	VersionID []IdentifierType `json:"VersionID,omitempty"`
}

// A class to define an exchange rate.
type ExchangeRateDetails struct {
	// The factor applied to the source currency to calculate the target currency.
	CalculationRate []NumericType `json:"CalculationRate,omitempty"`
	// The date on which the exchange rate was established.
	Date []DateType `json:"Date,omitempty"`
	// An identifier for the currency exchange market used as the source of this exchange rate.
	ExchangeMarketID []IdentifierType `json:"ExchangeMarketID,omitempty"`
	// A contract for foreign exchange.
	ForeignExchangeContract []ContractDetails `json:"ForeignExchangeContract,omitempty"`
	// A code signifying whether the calculation rate is a multiplier or a divisor.
	MathematicOperatorCode []CodeType `json:"MathematicOperatorCode,omitempty"`
	// In the case of a source currency with denominations of small value, the unit base.
	SourceCurrencyBaseRate []NumericType `json:"SourceCurrencyBaseRate,omitempty"`
	// The reference currency for this exchange rate; the currency from which the exchange is
	// being made.
	SourceCurrencyCode []CodeType `json:"SourceCurrencyCode"`
	// In the case of a target currency with denominations of small value, the unit base.
	TargetCurrencyBaseRate []NumericType `json:"TargetCurrencyBaseRate,omitempty"`
	// The target currency for this exchange rate; the currency to which the exchange is being
	// made.
	TargetCurrencyCode []CodeType `json:"TargetCurrencyCode"`
}

// A class to describe a set of payment terms.
type PaymentTermsDetails struct {
	// The monetary amount covered by these payment terms.
	Amount []AmountType `json:"Amount,omitempty"`
	// The currency exchange rate for purposes of these payment terms.
	ExchangeRate []ExchangeRateDetails `json:"ExchangeRate,omitempty"`
	// An identifier for this set of payment terms.
	ID []IdentifierType `json:"ID,omitempty"`
	// The due date for an installment payment for these payment terms.
	InstallmentDueDate []DateType `json:"InstallmentDueDate,omitempty"`
	// A reference to the payment terms used by the invoicing party. This may have been
	// requested of the payer by the payee to accompany its remittance.
	InvoicingPartyReference []TextType `json:"InvoicingPartyReference,omitempty"`
	// Free-form text conveying information that is not contained explicitly in other structures.
	Note []TextType `json:"Note,omitempty"`
	// The due date for these payment terms.
	PaymentDueDate []DateType `json:"PaymentDueDate,omitempty"`
	// An identifier for a means of payment associated with these payment terms.
	PaymentMeansID []IdentifierType `json:"PaymentMeansID,omitempty"`
	// The part of a payment, expressed as a percent, relevant for these payment terms.
	PaymentPercent []NumericType `json:"PaymentPercent,omitempty"`
	// The Uniform Resource Identifier (URI) of a document providing additional details
	// regarding these payment terms.
	PaymentTermsDetailsURI []IdentifierType `json:"PaymentTermsDetailsURI,omitempty"`
	// The monetary amount of the penalty for payment after the settlement period.
	PenaltyAmount []AmountType `json:"PenaltyAmount,omitempty"`
	// The period during which penalties may apply.
	PenaltyPeriod []PeriodDetails `json:"PenaltyPeriod,omitempty"`
	// The penalty for payment after the settlement period, expressed as a percentage of the
	// payment.
	PenaltySurchargePercent []NumericType `json:"PenaltySurchargePercent,omitempty"`
	// An identifier for a reference to a prepaid payment.
	PrepaidPaymentReferenceID []IdentifierType `json:"PrepaidPaymentReferenceID,omitempty"`
	// A code signifying the event during which these terms are offered.
	ReferenceEventCode []CodeType `json:"ReferenceEventCode,omitempty"`
	// The amount of a settlement discount offered for payment under these payment terms.
	SettlementDiscountAmount []AmountType `json:"SettlementDiscountAmount,omitempty"`
	// The percentage for the settlement discount that is offered for payment under these
	// payment terms.
	SettlementDiscountPercent []NumericType `json:"SettlementDiscountPercent,omitempty"`
	// The period during which settlement may occur.
	SettlementPeriod []PeriodDetails `json:"SettlementPeriod,omitempty"`
	// The period during which these payment terms are valid.
	ValidityPeriod []PeriodDetails `json:"ValidityPeriod,omitempty"`
}

// A class to describe an identifiable collection of one or more goods items to be
// transported between the consignor and the consignee. This information may be defined
// within a transport contract. A consignment may comprise more than one shipment (e.g.,
// when consolidated by a freight forwarder).
type ConsignmentDetails struct {
	// An indication that the transported goods in this consignment are animal foodstuffs (true)
	// or not (false).
	AnimalFoodIndicator []IndicatorType `json:"AnimalFoodIndicator,omitempty"`
	// The party holding the bill of lading for this consignment.
	BillOfLadingHolderParty []PartyDetails `json:"BillOfLadingHolderParty,omitempty"`
	// An identifier for this consignment, assigned by the broker.
	BrokerAssignedID []IdentifierType `json:"BrokerAssignedID,omitempty"`
	// An indication that the transported goods in this consignment are bulk cargoes (true) or
	// not (false).
	BulkCargoIndicator []IndicatorType `json:"BulkCargoIndicator,omitempty"`
	// An identifier for this consignment, assigned by the carrier.
	CarrierAssignedID []IdentifierType `json:"CarrierAssignedID,omitempty"`
	// The party providing the transport of goods in this consignment between named points.
	CarrierParty []PartyDetails `json:"CarrierParty,omitempty"`
	// Service instructions to the carrier, expressed as text.
	CarrierServiceInstructions []TextType `json:"CarrierServiceInstructions,omitempty"`
	// The weight upon which a charge is to be based.
	ChargeableWeightMeasure []MeasureType `json:"ChargeableWeightMeasure,omitempty"`
	// One of the child consignments of which a consolidated consignment is composed.
	ChildConsignment []ConsignmentDetails `json:"ChildConsignment,omitempty"`
	// The quantity of (consolidated) child consignments
	ChildConsignmentQuantity []QuantityType `json:"ChildConsignmentQuantity,omitempty"`
	// The terms of payment that apply to the collection of this consignment.
	CollectPaymentTerms []PaymentTermsDetails `json:"CollectPaymentTerms,omitempty"`
	// An identifier for this consignment, assigned by the consignee.
	ConsigneeAssignedID []IdentifierType `json:"ConsigneeAssignedID,omitempty"`
	// A party to which goods are consigned.
	ConsigneeParty []PartyDetails `json:"ConsigneeParty,omitempty"`
	// The count in this consignment considering goods items, child consignments, shipments
	ConsignmentQuantity []QuantityType `json:"ConsignmentQuantity,omitempty"`
	// An identifier for this consignment, assigned by the consignor.
	ConsignorAssignedID []IdentifierType `json:"ConsignorAssignedID,omitempty"`
	// The party consigning goods, as stipulated in the transport contract by the party ordering
	// transport.
	ConsignorParty []PartyDetails `json:"ConsignorParty,omitempty"`
	// An indicator that this consignment can be consolidated (true) or not (false).
	ConsolidatableIndicator []IndicatorType `json:"ConsolidatableIndicator,omitempty"`
	// A consolidated shipment (a shipment created by an act of consolidation).
	ConsolidatedShipment []ShipmentDetails `json:"ConsolidatedShipment,omitempty"`
	// An indication that the transported goods in this consignment are containerized cargoes
	// (true) or not (false).
	ContainerizedIndicator []IndicatorType `json:"ContainerizedIndicator,omitempty"`
	// An identifier for this consignment, assigned by the contracted carrier.
	ContractedCarrierAssignedID []IdentifierType `json:"ContractedCarrierAssignedID,omitempty"`
	// Service instructions for customs clearance, expressed as text.
	CustomsClearanceServiceInstructions []TextType `json:"CustomsClearanceServiceInstructions,omitempty"`
	// A class describing identifiers or references relating to customs procedures.
	CustomsDeclaration []CustomsDeclarationDetails `json:"CustomsDeclaration,omitempty"`
	// The total declared value for customs purposes of all the goods in this consignment,
	// regardless of whether they are subject to the same customs procedure, tariff/statistical
	// categorization, country information, or duty regime.
	DeclaredCustomsValueAmount []AmountType `json:"DeclaredCustomsValueAmount,omitempty"`
	// The value of this consignment, declared by the shipper or his agent solely for the
	// purpose of varying the carrier's level of liability from that provided in the contract of
	// carriage, in case of loss or damage to goods or delayed delivery.
	DeclaredForCarriageValueAmount []AmountType `json:"DeclaredForCarriageValueAmount,omitempty"`
	// The value, declared for statistical purposes, of those goods in this consignment that
	// have the same statistical heading.
	DeclaredStatisticsValueAmount []AmountType `json:"DeclaredStatisticsValueAmount,omitempty"`
	// A set of delivery instructions relating to this consignment.
	DeliveryInstructions []TextType `json:"DeliveryInstructions,omitempty"`
	// The conditions agreed upon between a seller and a buyer with regard to the delivery of
	// goods and/or services (e.g., CIF, FOB, or EXW from the INCOTERMS Terms of Delivery).
	DeliveryTerms []DeliveryTermsDetails `json:"DeliveryTerms,omitempty"`
	// The terms of payment for disbursement.
	DisbursementPaymentTerms []PaymentTermsDetails `json:"DisbursementPaymentTerms,omitempty"`
	// The party that makes the export declaration, or on behalf of which the export declaration
	// is made, and that is the owner of the goods in this consignment or has similar right of
	// disposal over them at the time when the declaration is accepted.
	ExporterParty []PartyDetails `json:"ExporterParty,omitempty"`
	// A charge for extra allowance.
	ExtraAllowanceCharge []AllowanceChargeDetails `json:"ExtraAllowanceCharge,omitempty"`
	// The final delivery party for this consignment.
	FinalDeliveryParty []PartyDetails `json:"FinalDeliveryParty,omitempty"`
	// The service for delivery to the consignee under the transport contract for this
	// consignment.
	FinalDeliveryTransportationService []TransportationServiceDetails `json:"FinalDeliveryTransportationService,omitempty"`
	// The country in which the goods in this consignment are to be delivered to the final
	// consignee or buyer.
	FinalDestinationCountry []CountryDetails `json:"FinalDestinationCountry,omitempty"`
	// The first arrival location in a transport. This would be a port for sea, an airport for
	// air, a terminal for rail, or a border post for land crossing.
	FirstArrivalPortLocation []LocationDetails `json:"FirstArrivalPortLocation,omitempty"`
	// Service instructions for the forwarder, expressed as text.
	ForwarderServiceInstructions []TextType `json:"ForwarderServiceInstructions,omitempty"`
	// The monetary amount that has to be or has been paid as calculated under the applicable
	// trade delivery.
	FreeOnBoardValueAmount []AmountType `json:"FreeOnBoardValueAmount,omitempty"`
	// A cost incurred by the shipper in moving goods, by whatever means, from one place to
	// another under the terms of the contract of carriage for this consignment. In addition to
	// transport costs, this may include such elements as packing, documentation, loading,
	// unloading, and insurance to the extent that they relate to the freight costs.
	FreightAllowanceCharge []AllowanceChargeDetails `json:"FreightAllowanceCharge,omitempty"`
	// An identifier for this consignment, assigned by the freight forwarder.
	FreightForwarderAssignedID []IdentifierType `json:"FreightForwarderAssignedID,omitempty"`
	// The party combining individual smaller consignments into a single larger shipment (the
	// consolidated shipment), which is sent to a counterpart that mirrors the consolidator's
	// activity by dividing the consolidated consignment into its original components.
	FreightForwarderParty []PartyDetails `json:"FreightForwarderParty,omitempty"`
	// An indication that the transported goods in this consignment are general cargoes (true)
	// or not (false).
	GeneralCargoIndicator []IndicatorType `json:"GeneralCargoIndicator,omitempty"`
	// The total volume of the goods referred to as one consignment.
	GrossVolumeMeasure []MeasureType `json:"GrossVolumeMeasure,omitempty"`
	// The total declared weight of the goods in this consignment, including packaging but
	// excluding the carrier's equipment.
	GrossWeightMeasure []MeasureType `json:"GrossWeightMeasure,omitempty"`
	// The handling required for this consignment, expressed as a code.
	HandlingCode []CodeType `json:"HandlingCode,omitempty"`
	// The handling required for this consignment, expressed as text.
	HandlingInstructions []TextType `json:"HandlingInstructions,omitempty"`
	// Instructions regarding haulage of this consignment, expressed as text.
	HaulageInstructions []TextType `json:"HaulageInstructions,omitempty"`
	// The party that would be notified of a hazardous item in this consignment.
	HazardousItemNotificationParty []PartyDetails `json:"HazardousItemNotificationParty,omitempty"`
	// An indication that the transported goods in this consignment are subject to an
	// international regulation concerning the carriage of dangerous goods (true) or not (false).
	HazardousRiskIndicator []IndicatorType `json:"HazardousRiskIndicator,omitempty"`
	// An indication that the transported goods in this consignment are for human consumption
	// (true) or not (false).
	HumanFoodIndicator []IndicatorType `json:"HumanFoodIndicator,omitempty"`
	// An identifier assigned to a collection of goods for both import and export.
	ID []IdentifierType `json:"ID"`
	// The party that makes an import declaration regarding this consignment, or on behalf of
	// which a customs clearing agent or other authorized person makes an import declaration
	// regarding this consignment. This may include a person who has possession of the goods or
	// to whom the goods are consigned.
	ImporterParty []PartyDetails `json:"ImporterParty,omitempty"`
	// Free-form text pertinent to this consignment, conveying information that is not contained
	// explicitly in other structures.
	Information []TextType `json:"Information,omitempty"`
	// The party holding the insurance for this consignment.
	InsuranceParty []PartyDetails `json:"InsuranceParty,omitempty"`
	// The amount of the premium payable to an insurance company for insuring the goods
	// contained in this consignment.
	InsurancePremiumAmount []AmountType `json:"InsurancePremiumAmount,omitempty"`
	// The amount covered by insurance for this consignment.
	InsuranceValueAmount []AmountType `json:"InsuranceValueAmount,omitempty"`
	// The final exporting location in a transport. This would be a port for sea, an airport for
	// air, a terminal for rail, or a border post for land crossing.
	LastExitPortLocation []LocationDetails `json:"LastExitPortLocation,omitempty"`
	// An indication that the transported goods are livestock (true) or not (false).
	LivestockIndicator []IndicatorType `json:"LivestockIndicator,omitempty"`
	// The total length in a means of transport or a piece of transport equipment which, given
	// the width and height of the transport means, will accommodate all of the consignments in
	// a single consolidation.
	LoadingLengthMeasure []MeasureType `json:"LoadingLengthMeasure,omitempty"`
	// An identifier for the loading sequence of this consignment.
	LoadingSequenceID []IdentifierType `json:"LoadingSequenceID,omitempty"`
	// The logistics operator party for this consignment.
	LogisticsOperatorParty []PartyDetails `json:"LogisticsOperatorParty,omitempty"`
	// A shipment stage during main carriage.
	MainCarriageShipmentStage []ShipmentStageDetails `json:"MainCarriageShipmentStage,omitempty"`
	// The party holding the mortgage for this consignment.
	MortgageHolderParty []PartyDetails `json:"MortgageHolderParty,omitempty"`
	// The total net weight of the goods in this consignment, exclusive of packaging.
	NetNetWeightMeasure []MeasureType `json:"NetNetWeightMeasure,omitempty"`
	// The total net volume of all goods items referred to as one consignment.
	NetVolumeMeasure []MeasureType `json:"NetVolumeMeasure,omitempty"`
	// The total net weight of all the goods items referred to as one consignment.
	NetWeightMeasure []MeasureType `json:"NetWeightMeasure,omitempty"`
	// The party to be notified upon arrival of goods and when special occurrences (usually
	// pre-defined) take place during a transportation service.
	NotifyParty []PartyDetails `json:"NotifyParty,omitempty"`
	// A shipment stage during on-carriage (usually refers to movement activity that takes place
	// after the container is discharged at a port of discharge).
	OnCarriageShipmentStage []ShipmentStageDetails `json:"OnCarriageShipmentStage,omitempty"`
	// The country from which the goods in this consignment were originally exported, without
	// any commercial transaction taking place in intermediate countries.
	OriginalDepartureCountry []CountryDetails `json:"OriginalDepartureCountry,omitempty"`
	// The original despatch (sender) party for this consignment.
	OriginalDespatchParty []PartyDetails `json:"OriginalDespatchParty,omitempty"`
	// The service for pickup from the consignor under the transport contract for this
	// consignment.
	OriginalDespatchTransportationService []TransportationServiceDetails `json:"OriginalDespatchTransportationService,omitempty"`
	// The terms of payment between the parties (such as logistics service client, logistics
	// service provider) in a transaction.
	PaymentTerms []PaymentTermsDetails `json:"PaymentTerms,omitempty"`
	// An identifier for this consignment, assigned by the performing carrier.
	PerformingCarrierAssignedID []IdentifierType `json:"PerformingCarrierAssignedID,omitempty"`
	// The party performing the carriage of this consignment.
	PerformingCarrierParty []PartyDetails `json:"PerformingCarrierParty,omitempty"`
	// The delivery of this consignment planned by the party responsible for providing the
	// transportation service (the transport service provider).
	PlannedDeliveryTransportEvent []TransportEventDetails `json:"PlannedDeliveryTransportEvent,omitempty"`
	// The pickup of this consignment planned by the party responsible for providing the
	// transportation service (the transport service provider).
	PlannedPickupTransportEvent []TransportEventDetails `json:"PlannedPickupTransportEvent,omitempty"`
	// A shipment stage during precarriage (usually refers to movement activity that takes place
	// prior to the container being loaded at a port of loading).
	PreCarriageShipmentStage []ShipmentStageDetails `json:"PreCarriageShipmentStage,omitempty"`
	// The terms of payment for prepayment.
	PrepaidPaymentTerms []PaymentTermsDetails `json:"PrepaidPaymentTerms,omitempty"`
	// Remarks concerning the complete consignment, to be printed on the transport document.
	Remarks []TextType `json:"Remarks,omitempty"`
	// The delivery of this consignment requested by the party requesting a transportation
	// service (the transport user).
	RequestedDeliveryTransportEvent []TransportEventDetails `json:"RequestedDeliveryTransportEvent,omitempty"`
	// The pickup of this consignment requested by the party requesting a transportation service
	// (the transport user).
	RequestedPickupTransportEvent []TransportEventDetails `json:"RequestedPickupTransportEvent,omitempty"`
	// A sequence identifier for this consignment.
	SequenceID []IdentifierType `json:"SequenceID,omitempty"`
	// A code signifying the priority or level of service required for this consignment.
	ShippingPriorityLevelCode []CodeType `json:"ShippingPriorityLevelCode,omitempty"`
	// Special instructions relating to this consignment.
	SpecialInstructions []TextType `json:"SpecialInstructions,omitempty"`
	// An indication that the transported goods in this consignment require special security
	// (true) or not (false).
	SpecialSecurityIndicator []IndicatorType `json:"SpecialSecurityIndicator,omitempty"`
	// Special service instructions, expressed as text.
	SpecialServiceInstructions []TextType `json:"SpecialServiceInstructions,omitempty"`
	// An indicator that this consignment has been split in transit (true) or not (false).
	SplitConsignmentIndicator []IndicatorType `json:"SplitConsignmentIndicator,omitempty"`
	// The status of a particular condition associated with this consignment.
	Status []StatusDetails `json:"Status,omitempty"`
	// A substitute party performing the carriage of this consignment.
	SubstituteCarrierParty []PartyDetails `json:"SubstituteCarrierParty,omitempty"`
	// A textual summary description of the consignment.
	SummaryDescription []TextType `json:"SummaryDescription,omitempty"`
	// A code signifying the tariff applied to this consignment.
	TariffCode []CodeType `json:"TariffCode,omitempty"`
	// Text describing the tariff applied to this consignment.
	TariffDescription []TextType `json:"TariffDescription,omitempty"`
	// An indication that this consignment will be paid for by a third party (true) or not
	// (false).
	ThirdPartyPayerIndicator []IndicatorType `json:"ThirdPartyPayerIndicator,omitempty"`
	// The total number of goods items in this consignment.
	TotalGoodsItemQuantity []QuantityType `json:"TotalGoodsItemQuantity,omitempty"`
	// The total of all invoice amounts declared in this consignment.
	TotalInvoiceAmount []AmountType `json:"TotalInvoiceAmount,omitempty"`
	// The total number of packages associated with a Consignment.
	TotalPackagesQuantity []QuantityType `json:"TotalPackagesQuantity,omitempty"`
	// The number of pieces of transport handling equipment (pallets, boxes, cases, etc.) in
	// this consignment.
	TotalTransportHandlingUnitQuantity []QuantityType `json:"TotalTransportHandlingUnitQuantity,omitempty"`
	// One of the countries through which goods or passengers in this consignment are routed
	// between the country of original departure and the country of final destination.
	TransitCountry []CountryDetails `json:"TransitCountry,omitempty"`
	// The party providing transport advice this consignment.
	TransportAdvisorParty []PartyDetails `json:"TransportAdvisorParty,omitempty"`
	// A transport contract relating to this consignment.
	TransportContract []ContractDetails `json:"TransportContract,omitempty"`
	// A class describing a significant occurrence or happening related to the transportation of
	// goods.
	TransportEvent []TransportEventDetails `json:"TransportEvent,omitempty"`
	// A transport handling unit used for loose and containerized goods.
	TransportHandlingUnit []TransportHandlingUnitDetails `json:"TransportHandlingUnit,omitempty"`
}

// A class defining an identifiable collection of one or more goods items to be transported
// between the seller party and the buyer party. This information may be defined within a
// commercial contract. A shipment can be transported in different consignments (e.g., split
// for logistical purposes).
type ShipmentDetails struct {
	// A consignment covering this shipment.
	Consignment []ConsignmentDetails `json:"Consignment,omitempty"`
	// The total number of consignments within this shipment.
	ConsignmentQuantity []QuantityType `json:"ConsignmentQuantity,omitempty"`
	// The total declared value for customs purposes of those goods in this shipment that are
	// subject to the same customs procedure and have the same tariff/statistical heading,
	// country information, and duty regime.
	DeclaredCustomsValueAmount []AmountType `json:"DeclaredCustomsValueAmount,omitempty"`
	// The value of this shipment, declared by the shipper or his agent solely for the purpose
	// of varying the carrier's level of liability from that provided in the contract of
	// carriage, in case of loss or damage to goods or delayed delivery.
	DeclaredForCarriageValueAmount []AmountType `json:"DeclaredForCarriageValueAmount,omitempty"`
	// The value, declared for statistical purposes, of those goods in this shipment that have
	// the same statistical heading.
	DeclaredStatisticsValueAmount []AmountType `json:"DeclaredStatisticsValueAmount,omitempty"`
	// The delivery of this shipment.
	Delivery []DeliveryDetails `json:"Delivery,omitempty"`
	// Delivery instructions relating to this shipment.
	DeliveryInstructions []TextType `json:"DeliveryInstructions,omitempty"`
	// The country from which the goods were originally exported, without any commercial
	// transaction taking place in intermediate countries.
	ExportCountry []CountryDetails `json:"ExportCountry,omitempty"`
	// The first arrival location of a shipment. This would be a port for sea, an airport for
	// air, a terminal for rail, or a border post for land crossing.
	FirstArrivalPortLocation []LocationDetails `json:"FirstArrivalPortLocation,omitempty"`
	// The monetary amount that has to be or has been paid as calculated under the applicable
	// trade delivery.
	FreeOnBoardValueAmount []AmountType `json:"FreeOnBoardValueAmount,omitempty"`
	// A cost incurred by the shipper in moving goods, by whatever means, from one place to
	// another under the terms of the contract of carriage. In addition to transport costs, this
	// may include such elements as packing, documentation, loading, unloading, and insurance to
	// the extent that they relate to the freight costs.
	FreightAllowanceCharge []AllowanceChargeDetails `json:"FreightAllowanceCharge,omitempty"`
	// A goods item included in this shipment.
	GoodsItem []GoodsItemDetails `json:"GoodsItem,omitempty"`
	// The total volume of the goods in this shipment, including packaging.
	GrossVolumeMeasure []MeasureType `json:"GrossVolumeMeasure,omitempty"`
	// The total gross weight of a shipment; the weight of the goods plus packaging plus
	// transport equipment.
	GrossWeightMeasure []MeasureType `json:"GrossWeightMeasure,omitempty"`
	// The handling required for this shipment, expressed as a code.
	HandlingCode []CodeType `json:"HandlingCode,omitempty"`
	// The handling required for this shipment, expressed as text.
	HandlingInstructions []TextType `json:"HandlingInstructions,omitempty"`
	// An identifier for this shipment.
	ID []IdentifierType `json:"ID"`
	// Free-form text pertinent to this shipment, conveying information that is not contained
	// explicitly in other structures.
	Information []TextType `json:"Information,omitempty"`
	// The amount covered by insurance for this shipment.
	InsuranceValueAmount []AmountType `json:"InsuranceValueAmount,omitempty"`
	// The final exporting location for a shipment. This would be a port for sea, an airport for
	// air, a terminal for rail, or a border post for land crossing.
	LastExitPortLocation []LocationDetails `json:"LastExitPortLocation,omitempty"`
	// The total net weight of this shipment, excluding packaging and transport equipment.
	NetNetWeightMeasure []MeasureType `json:"NetNetWeightMeasure,omitempty"`
	// The total volume of the goods in this shipment, excluding packaging and transport
	// equipment.
	NetVolumeMeasure []MeasureType `json:"NetVolumeMeasure,omitempty"`
	// The net weight of this shipment, excluding packaging.
	NetWeightMeasure []MeasureType `json:"NetWeightMeasure,omitempty"`
	// The region in which the goods have been produced or manufactured, according to criteria
	// laid down for the purposes of application of the customs tariff, or of quantitative
	// restrictions, or of any other measure related to trade.
	OriginAddress []AddressDetails `json:"OriginAddress,omitempty"`
	// The address to which a shipment should be returned.
	ReturnAddress []AddressDetails `json:"ReturnAddress,omitempty"`
	// A stage in the transport movement of this shipment.
	ShipmentStage []ShipmentStageDetails `json:"ShipmentStage,omitempty"`
	// A code signifying the priority or level of service required for this shipment.
	ShippingPriorityLevelCode []CodeType `json:"ShippingPriorityLevelCode,omitempty"`
	// Special instructions relating to this shipment.
	SpecialInstructions []TextType `json:"SpecialInstructions,omitempty"`
	// An indicator that the consignment has been split in transit (true) or not (false).
	SplitConsignmentIndicator []IndicatorType `json:"SplitConsignmentIndicator,omitempty"`
	// The total number of goods items in this shipment.
	TotalGoodsItemQuantity []QuantityType `json:"TotalGoodsItemQuantity,omitempty"`
	// The number of pieces of transport handling equipment (pallets, boxes, cases, etc.) in
	// this shipment.
	TotalTransportHandlingUnitQuantity []QuantityType `json:"TotalTransportHandlingUnitQuantity,omitempty"`
	// A transport handling unit associated with this shipment.
	TransportHandlingUnit []TransportHandlingUnitDetails `json:"TransportHandlingUnit,omitempty"`
}

// A class to describe a delivery.
type DeliveryDetails struct {
	// The actual date of delivery.
	ActualDeliveryDate []DateType `json:"ActualDeliveryDate,omitempty"`
	// The actual time of delivery.
	ActualDeliveryTime []TimeType `json:"ActualDeliveryTime,omitempty"`
	// An alternative delivery location.
	AlternativeDeliveryLocation []LocationDetails `json:"AlternativeDeliveryLocation,omitempty"`
	// The party responsible for delivering the goods.
	CarrierParty []PartyDetails `json:"CarrierParty,omitempty"`
	// The delivery address.
	DeliveryAddress []AddressDetails `json:"DeliveryAddress,omitempty"`
	// The delivery location.
	DeliveryLocation []LocationDetails `json:"DeliveryLocation,omitempty"`
	// The party to whom the goods are delivered.
	DeliveryParty []PartyDetails `json:"DeliveryParty,omitempty"`
	// Terms and conditions relating to the delivery.
	DeliveryTerms []DeliveryTermsDetails `json:"DeliveryTerms,omitempty"`
	// The despatch (pickup) associated with this delivery.
	Despatch []DespatchDetails `json:"Despatch,omitempty"`
	// The period estimated for delivery.
	EstimatedDeliveryPeriod []PeriodDetails `json:"EstimatedDeliveryPeriod,omitempty"`
	// An identifier for this delivery.
	ID []IdentifierType `json:"ID,omitempty"`
	// The latest date of delivery allowed by the buyer.
	LatestDeliveryDate []DateType `json:"LatestDeliveryDate,omitempty"`
	// The latest time of delivery allowed by the buyer.
	LatestDeliveryTime []TimeType `json:"LatestDeliveryTime,omitempty"`
	// The maximum delivery unit for this delivery.
	MaximumDeliveryUnit []DeliveryUnitDetails `json:"MaximumDeliveryUnit,omitempty"`
	// The maximum quantity of items, child consignments, shipments in this delivery.
	MaximumQuantity []QuantityType `json:"MaximumQuantity,omitempty"`
	// The minimum delivery unit for this delivery.
	MinimumDeliveryUnit []DeliveryUnitDetails `json:"MinimumDeliveryUnit,omitempty"`
	// The minimum quantity of items, child consignments, shipments in this delivery.
	MinimumQuantity []QuantityType `json:"MinimumQuantity,omitempty"`
	// A party to be notified of this delivery.
	NotifyParty []PartyDetails `json:"NotifyParty,omitempty"`
	// The period promised for delivery.
	PromisedDeliveryPeriod []PeriodDetails `json:"PromisedDeliveryPeriod,omitempty"`
	// The quantity of items, child consignments, shipments in this delivery.
	Quantity []QuantityType `json:"Quantity,omitempty"`
	// An identifier used for approval of access to delivery locations (e.g., port terminals).
	ReleaseID []IdentifierType `json:"ReleaseID,omitempty"`
	// The period requested for delivery.
	RequestedDeliveryPeriod []PeriodDetails `json:"RequestedDeliveryPeriod,omitempty"`
	// The shipment being delivered.
	Shipment []ShipmentDetails `json:"Shipment,omitempty"`
	// The delivery Tracking ID (for transport tracking).
	TrackingID []IdentifierType `json:"TrackingID,omitempty"`
}

// A class to define a reference to a line in a document.
type LineReferenceDetails struct {
	// A reference to the document containing the referenced line.
	DocumentReference []DocumentReferenceDetails `json:"DocumentReference,omitempty"`
	// Identifies the referenced line in the document.
	LineID []IdentifierType `json:"LineID"`
	// A code signifying the status of the referenced line with respect to its original state.
	LineStatusCode []CodeType `json:"LineStatusCode,omitempty"`
	// A universally unique identifier for this line reference.
	UUID []IdentifierType `json:"UUID,omitempty"`
}

// A class to describe an item of trade. It includes a generic description applicable to all
// examples of the item together with optional subsidiary descriptions of any number of
// actual instances of the type.
type ItemDetails struct {
	// Further details regarding this item (e.g., the URL of a relevant web page).
	AdditionalInformation []TextType `json:"AdditionalInformation,omitempty"`
	// An additional identifier for this item.
	AdditionalItemIdentification []ItemIdentificationDetails `json:"AdditionalItemIdentification,omitempty"`
	// An additional property of this item.
	AdditionalItemProperty []ItemPropertyDetails `json:"AdditionalItemProperty,omitempty"`
	// A brand name of this item.
	BrandName []TextType `json:"BrandName,omitempty"`
	// Identifying information for this item, assigned by the buyer.
	BuyersItemIdentification []ItemIdentificationDetails `json:"BuyersItemIdentification,omitempty"`
	// A reference to the catalogue in which this item appears.
	CatalogueDocumentReference []DocumentReferenceDetails `json:"CatalogueDocumentReference,omitempty"`
	// An indicator that this item was ordered from a catalogue (true) or not (false).
	CatalogueIndicator []IndicatorType `json:"CatalogueIndicator,omitempty"`
	// Identifying information for this item, assigned according to a cataloguing system.
	CatalogueItemIdentification []ItemIdentificationDetails `json:"CatalogueItemIdentification,omitempty"`
	// A certificate associated with this item.
	Certificate []CertificateDetails `json:"Certificate,omitempty"`
	// A tax category applicable to this item.
	ClassifiedTaxCategory []TaxCategoryDetails `json:"ClassifiedTaxCategory,omitempty"`
	// A classification of this item according to a specific system for classifying commodities.
	CommodityClassification []CommodityClassificationDetails `json:"CommodityClassification,omitempty"`
	// Text describing this item.
	Description []TextType `json:"Description,omitempty"`
	// One of the measurable dimensions (length, mass, weight, or volume) of this item.
	Dimension []DimensionDetails `json:"Dimension,omitempty"`
	// Information pertaining to this item as a hazardous item.
	HazardousItem []HazardousItemDetails `json:"HazardousItem,omitempty"`
	// An indication that the transported item, as delivered, is subject to an international
	// regulation concerning the carriage of dangerous goods (true) or not (false).
	HazardousRiskIndicator []IndicatorType `json:"HazardousRiskIndicator,omitempty"`
	// The party responsible for specification of this item.
	InformationContentProviderParty []PartyDetails `json:"InformationContentProviderParty,omitempty"`
	// A trackable, unique instantiation of this item.
	ItemInstance []ItemInstanceDetails `json:"ItemInstance,omitempty"`
	// A reference to a specification document for this item.
	ItemSpecificationDocumentReference []DocumentReferenceDetails `json:"ItemSpecificationDocumentReference,omitempty"`
	// A keyword (search string) for this item, assigned by the seller party. Can also be a
	// synonym for the name of the item.
	Keyword []TextType `json:"Keyword,omitempty"`
	// The manufacturer of this item.
	ManufacturerParty []PartyDetails `json:"ManufacturerParty,omitempty"`
	// Identifying information for this item, assigned by the manufacturer.
	ManufacturersItemIdentification []ItemIdentificationDetails `json:"ManufacturersItemIdentification,omitempty"`
	// A model name of this item.
	ModelName []TextType `json:"ModelName,omitempty"`
	// A short name optionally given to this item, such as a name from a catalogue, as distinct
	// from a description.
	Name []TextType `json:"Name,omitempty"`
	// A region (not country) of origin of this item.
	OriginAddress []AddressDetails `json:"OriginAddress,omitempty"`
	// The country of origin of this item.
	OriginCountry []CountryDetails `json:"OriginCountry,omitempty"`
	// The unit packaging quantity; the number of subunits making up this item.
	PackQuantity []QuantityType `json:"PackQuantity,omitempty"`
	// The number of items in a pack of this item.
	PackSizeNumeric []NumericType `json:"PackSizeNumeric,omitempty"`
	// Identifying information for this item, assigned by the seller.
	SellersItemIdentification []ItemIdentificationDetails `json:"SellersItemIdentification,omitempty"`
	// Identifying information for this item, assigned according to a standard system.
	StandardItemIdentification []ItemIdentificationDetails `json:"StandardItemIdentification,omitempty"`
	// A set of sales conditions applying to this item.
	TransactionConditions []TransactionConditionsDetails `json:"TransactionConditions,omitempty"`
}

// A class for assigning identifying information to an item.
type ItemIdentificationDetails struct {
	// An identifier for a system of barcodes.
	BarcodeSymbologyID []IdentifierType `json:"BarcodeSymbologyID,omitempty"`
	// An extended identifier for the item that identifies the item with specific properties,
	// e.g., Item 123 = Chair / Item 123 Ext 45 = brown chair. Two chairs can have the same item
	// number, but one is brown. The other is white.
	ExtendedID []IdentifierType `json:"ExtendedID,omitempty"`
	// An identifier for the item.
	ID []IdentifierType `json:"ID"`
	// The party that issued this item identification.
	IssuerParty []PartyDetails `json:"IssuerParty,omitempty"`
	// A measurable dimension (length, mass, weight, or volume) of the item.
	MeasurementDimension []DimensionDetails `json:"MeasurementDimension,omitempty"`
	// A physical attribute of the item.
	PhysicalAttribute []PhysicalAttributeDetails `json:"PhysicalAttribute,omitempty"`
}

// A class to define a measurable dimension (length, mass, weight, volume, or area) of an
// item.
type DimensionDetails struct {
	// An identifier for the attribute to which the measure applies.
	AttributeID []IdentifierType `json:"AttributeID"`
	// Text describing the measurement attribute.
	Description []TextType `json:"Description,omitempty"`
	// The maximum value in a range of measurement of this dimension.
	MaximumMeasure []MeasureType `json:"MaximumMeasure,omitempty"`
	// The measurement value.
	Measure []MeasureType `json:"Measure,omitempty"`
	// The minimum value in a range of measurement of this dimension.
	MinimumMeasure []MeasureType `json:"MinimumMeasure,omitempty"`
}

// A class to describe a physical attribute.
type PhysicalAttributeDetails struct {
	// An identifier for this physical attribute.
	AttributeID []IdentifierType `json:"AttributeID"`
	// A description of the physical attribute, expressed as text.
	Description []TextType `json:"Description,omitempty"`
	// A description of the physical attribute, expressed as a code.
	DescriptionCode []CodeType `json:"DescriptionCode,omitempty"`
	// A code signifying the position of this physical attribute.
	PositionCode []CodeType `json:"PositionCode,omitempty"`
}

// A class to describe a specific property of an item.
type ItemPropertyDetails struct {
	// An identifier for this property of an item.
	ID []IdentifierType `json:"ID,omitempty"`
	// A code signifying the importance of this property in using it to describe a related Item.
	ImportanceCode []CodeType `json:"ImportanceCode,omitempty"`
	// A description of the property group to which this item property belongs.
	ItemPropertyGroup []ItemPropertyGroupDetails `json:"ItemPropertyGroup,omitempty"`
	// A range of values for this item property.
	ItemPropertyRange []ItemPropertyRangeDetails `json:"ItemPropertyRange,omitempty"`
	// The value expressed as a text in case the property is a value in a list. For example, a
	// colour.
	ListValue []TextType `json:"ListValue,omitempty"`
	// The name of this item property.
	Name []TextType `json:"Name"`
	// The name of this item property, expressed as a code.
	NameCode []CodeType `json:"NameCode,omitempty"`
	// The range of values for the dimensions of this property.
	RangeDimension []DimensionDetails `json:"RangeDimension,omitempty"`
	// The method of testing the value of this item property.
	TestMethod []TextType `json:"TestMethod,omitempty"`
	// The period during which this item property is valid.
	UsabilityPeriod []PeriodDetails `json:"UsabilityPeriod,omitempty"`
	// The value of this item property, expressed as text.
	Value []TextType `json:"Value,omitempty"`
	// Text qualifying the value of the property.
	ValueQualifier []TextType `json:"ValueQualifier,omitempty"`
	// The value of this item property, expressed as a quantity.
	ValueQuantity []QuantityType `json:"ValueQuantity,omitempty"`
}

// A class to describe a property group or classification.
type ItemPropertyGroupDetails struct {
	// An identifier for this group of item properties.
	ID []IdentifierType `json:"ID"`
	// A code signifying the importance of this property group in using it to describe a
	// required Item.
	ImportanceCode []CodeType `json:"ImportanceCode,omitempty"`
	// The name of this item property group.
	Name []TextType `json:"Name,omitempty"`
}

// A class to describe a range of values for an item property.
type ItemPropertyRangeDetails struct {
	// The maximum value in this range of values.
	MaximumValue []TextType `json:"MaximumValue,omitempty"`
	// The minimum value in this range of values.
	MinimumValue []TextType `json:"MinimumValue,omitempty"`
}

// A class to define a certificate applied to the item. Certificated can be a requirement to
// sell goods or services in a jurisdiction.
type CertificateDetails struct {
	// The type of this certificate, expressed as a code. The type specifies what array it
	// belongs to, e.g.. Environmental, security, health improvement etc.
	CertificateType []TextType `json:"CertificateType"`
	// The type of this certificate, expressed as a code. The type specifies what array it
	// belongs to, e.g.. Environmental, security, health improvement etc.
	CertificateTypeCode []CodeType `json:"CertificateTypeCode"`
	// A reference to a document relevant to this certificate or an application for this
	// certificate.
	DocumentReference []DocumentReferenceDetails `json:"DocumentReference,omitempty"`
	// An identifier for this certificate.
	ID []IdentifierType `json:"ID"`
	// The authorized organization that issued this certificate, the provider of the certificate.
	IssuerParty []PartyDetails `json:"IssuerParty"`
	// Remarks by the applicant for this certificate.
	Remarks []TextType `json:"Remarks,omitempty"`
	// A signature applied to this certificate.
	Signature []SignatureDetails `json:"Signature,omitempty"`
}

// A class to define a signature.
type SignatureDetails struct {
	// The method used to perform XML canonicalization of this signature.
	CanonicalizationMethod []TextType `json:"CanonicalizationMethod,omitempty"`
	// The actual encoded signature (e.g., in XMLDsig format).
	DigitalSignatureAttachment []AttachmentDetails `json:"DigitalSignatureAttachment,omitempty"`
	// An identifier for this signature.
	ID []IdentifierType `json:"ID"`
	// Free-form text conveying information that is not contained explicitly in other
	// structures; in particular, information regarding the circumstances in which the signature
	// is being used.
	Note []TextType `json:"Note,omitempty"`
	// A reference to the document that the signature applies to. For evidentiary purposes, this
	// may be the document image that the signatory party saw when applying their signature.
	OriginalDocumentReference []DocumentReferenceDetails `json:"OriginalDocumentReference,omitempty"`
	// The signing party.
	SignatoryParty []PartyDetails `json:"SignatoryParty,omitempty"`
	// Text describing the method of signature.
	SignatureMethod []TextType `json:"SignatureMethod,omitempty"`
	// The date upon which this signature was verified.
	ValidationDate []DateType `json:"ValidationDate,omitempty"`
	// The time at which this signature was verified.
	ValidationTime []TimeType `json:"ValidationTime,omitempty"`
	// An identifier for the organization, person, service, or server that verified this
	// signature.
	ValidatorID []IdentifierType `json:"ValidatorID,omitempty"`
}

// A class to describe the classification of a commodity.
type CommodityClassificationDetails struct {
	// A mutually agreed code signifying the type of cargo for purposes of commodity
	// classification.
	CargoTypeCode []CodeType `json:"CargoTypeCode,omitempty"`
	// The harmonized international commodity code for cross border and regulatory (customs and
	// trade statistics) purposes.
	CommodityCode []CodeType `json:"CommodityCode,omitempty"`
	// A code signifying the trade classification of the commodity.
	ItemClassificationCode []CodeType `json:"ItemClassificationCode,omitempty"`
	// A code defined by a specific maintenance agency signifying the high-level nature of the
	// commodity.
	NatureCode []CodeType `json:"NatureCode,omitempty"`
}

// A class to describe a hazardous item.
type HazardousItemDetails struct {
	// Text providing further information about the hazardous substance.
	AdditionalInformation []TextType `json:"AdditionalInformation,omitempty"`
	// Another temperature relevant to the handling of this hazardous item.
	AdditionalTemperature []TemperatureDetails `json:"AdditionalTemperature,omitempty"`
	// The name of the category of hazard that applies to the Item.
	CategoryName []TextType `json:"CategoryName,omitempty"`
	// The individual, group, or body to be contacted in case of a hazardous incident associated
	// with this item.
	ContactParty []PartyDetails `json:"ContactParty,omitempty"`
	// A code signifying the emergency procedures for this hazardous item.
	EmergencyProceduresCode []CodeType `json:"EmergencyProceduresCode,omitempty"`
	// The threshold temperature at which emergency procedures apply in the handling of
	// temperature-controlled goods.
	EmergencyTemperature []TemperatureDetails `json:"EmergencyTemperature,omitempty"`
	// The flashpoint temperature of this hazardous item; i.e., the lowest temperature at which
	// vapors above a volatile combustible substance ignite in air when exposed to flame.
	FlashpointTemperature []TemperatureDetails `json:"FlashpointTemperature,omitempty"`
	// An identifier for the hazard class applicable to this hazardous item as defined by the
	// relevant regulation authority (e.g., the IMDG Class Number of the SOLAS Convention of IMO
	// and the ADR/RID Class Number for the road/rail environment).
	HazardClassID []IdentifierType `json:"HazardClassID,omitempty"`
	// A code signifying a kind of hazard for a material.
	HazardousCategoryCode []CodeType `json:"HazardousCategoryCode,omitempty"`
	// Information related to the transit of this kind of hazardous goods.
	HazardousGoodsTransit []HazardousGoodsTransitDetails `json:"HazardousGoodsTransit,omitempty"`
	// An identifier for this hazardous item.
	ID []IdentifierType `json:"ID,omitempty"`
	// The number for the lower part of the orange hazard placard required on the means of
	// transport.
	LowerOrangeHazardPlacardID []IdentifierType `json:"LowerOrangeHazardPlacardID,omitempty"`
	// An identifier to the marking of the Hazardous Item
	MarkingID []IdentifierType `json:"MarkingID,omitempty"`
	// A code signifying a medical first aid guide appropriate to this hazardous item.
	MedicalFirstAidGuideCode []CodeType `json:"MedicalFirstAidGuideCode,omitempty"`
	// The volume of this hazardous item, excluding packaging and transport equipment.
	NetVolumeMeasure []MeasureType `json:"NetVolumeMeasure,omitempty"`
	// The net weight of this hazardous item, excluding packaging.
	NetWeightMeasure []MeasureType `json:"NetWeightMeasure,omitempty"`
	// Text of the placard endorsement that is to be shown on the shipping papers for this
	// hazardous item. Can also be used for the number of the orange placard (lower part)
	// required on the means of transport.
	PlacardEndorsement []TextType `json:"PlacardEndorsement,omitempty"`
	// Text of the placard notation corresponding to the hazard class of this hazardous item.
	// Can also be the hazard identification number of the orange placard (upper part) required
	// on the means of transport.
	PlacardNotation []TextType `json:"PlacardNotation,omitempty"`
	// The quantity of goods items in this hazardous item that are hazardous.
	Quantity []QuantityType `json:"Quantity,omitempty"`
	// A secondary hazard associated with this hazardous item.
	SecondaryHazard []SecondaryHazardDetails `json:"SecondaryHazard,omitempty"`
	// The full technical name of a specific hazardous substance contained in this goods item.
	TechnicalName []TextType `json:"TechnicalName,omitempty"`
	// The UN code for this kind of hazardous item.
	UNDGCode []CodeType `json:"UNDGCode,omitempty"`
	// The number for the upper part of the orange hazard placard required on the means of
	// transport.
	UpperOrangeHazardPlacardID []IdentifierType `json:"UpperOrangeHazardPlacardID,omitempty"`
}

// A class to describe a measurement of temperature.
type TemperatureDetails struct {
	// An identifier for this temperature measurement.
	AttributeID []IdentifierType `json:"AttributeID"`
	// Text describing this temperature measurement.
	Description []TextType `json:"Description,omitempty"`
	// The value of this temperature measurement.
	Measure []MeasureType `json:"Measure"`
}

// A class to describe hazardous goods in transit.
type HazardousGoodsTransitDetails struct {
	// A code signifying the set of legal regulations governing the transportation of the
	// hazardous goods.
	HazardousRegulationCode []CodeType `json:"HazardousRegulationCode,omitempty"`
	// A code signifying the Inhalation Toxicity Hazard Zone for the hazardous goods, as defined
	// by the US Department of Transportation.
	InhalationToxicityZoneCode []CodeType `json:"InhalationToxicityZoneCode,omitempty"`
	// The maximum temperature at which the hazardous goods can safely be transported.
	MaximumTemperature []TemperatureDetails `json:"MaximumTemperature,omitempty"`
	// The minimum temperature at which the hazardous goods can safely be transported.
	MinimumTemperature []TemperatureDetails `json:"MinimumTemperature,omitempty"`
	// A code signifying the packaging requirement for transportation of the hazardous goods as
	// assigned by IATA, IMDB, ADR, RID etc.
	PackingCriteriaCode []CodeType `json:"PackingCriteriaCode,omitempty"`
	// A code signifying authorization for the transportation of hazardous cargo.
	TransportAuthorizationCode []CodeType `json:"TransportAuthorizationCode,omitempty"`
	// An identifier for a transport emergency card describing the actions to be taken in an
	// emergency in transporting the hazardous goods. It may be the identity number of a
	// hazardous emergency response plan assigned by the appropriate authority.
	TransportEmergencyCardCode []CodeType `json:"TransportEmergencyCardCode,omitempty"`
}

// A class to describe a secondary hazard associated with a hazardous item.
type SecondaryHazardDetails struct {
	// A code signifying the emergency procedures for this secondary hazard.
	EmergencyProceduresCode []CodeType `json:"EmergencyProceduresCode,omitempty"`
	// Additional information about the hazardous substance, which can be used (for example) to
	// specify the type of regulatory requirements that apply to this secondary hazard.
	Extension []TextType `json:"Extension,omitempty"`
	// An identifier for this secondary hazard.
	ID []IdentifierType `json:"ID,omitempty"`
	// Text of the placard endorsement for this secondary hazard that is to be shown on the
	// shipping papers for a hazardous item. Can also be used for the number of the orange
	// placard (lower part) required on the means of transport.
	PlacardEndorsement []TextType `json:"PlacardEndorsement,omitempty"`
	// Text of the placard notation corresponding to the hazard class of this secondary hazard.
	// Can also be the hazard identification number of the orange placard (upper part) required
	// on the means of transport.
	PlacardNotation []TextType `json:"PlacardNotation,omitempty"`
}

// A class to describe a specific, trackable instance of an item.
type ItemInstanceDetails struct {
	// An additional property of this item instance.
	AdditionalItemProperty []ItemPropertyDetails `json:"AdditionalItemProperty,omitempty"`
	// The date before which it is best to use this item instance.
	BestBeforeDate []DateType `json:"BestBeforeDate,omitempty"`
	// The lot identifier of this item instance (the identifier that allows recall of the item
	// if necessary).
	LotIdentification []LotIdentificationDetails `json:"LotIdentification,omitempty"`
	// The date on which this item instance was manufactured.
	ManufactureDate []DateType `json:"ManufactureDate,omitempty"`
	// The time at which this item instance was manufactured.
	ManufactureTime []TimeType `json:"ManufactureTime,omitempty"`
	// An identifier used for tracing this item instance, such as the EPC number used in RFID.
	ProductTraceID []IdentifierType `json:"ProductTraceID,omitempty"`
	// The registration identifier of this item instance.
	RegistrationID []IdentifierType `json:"RegistrationID,omitempty"`
	// The serial number of this item instance.
	SerialID []IdentifierType `json:"SerialID,omitempty"`
}

// A class for defining a lot identifier (the identifier of a set of item instances that
// would be used in case of a recall of that item).
type LotIdentificationDetails struct {
	// An additional property of the lot.
	AdditionalItemProperty []ItemPropertyDetails `json:"AdditionalItemProperty,omitempty"`
	// The expiry date of the lot.
	ExpiryDate []DateType `json:"ExpiryDate,omitempty"`
	// An identifier for the lot.
	LotNumberID []IdentifierType `json:"LotNumberID,omitempty"`
}

// A class to describe purchasing, sales, or payment conditions.
type TransactionConditionsDetails struct {
	// A code signifying a type of action relating to sales or payment conditions.
	ActionCode []CodeType `json:"ActionCode,omitempty"`
	// Text describing the transaction conditions.
	Description []TextType `json:"Description,omitempty"`
	// A document associated with these transaction conditions.
	DocumentReference []DocumentReferenceDetails `json:"DocumentReference,omitempty"`
	// An identifier for conditions of the transaction, typically purchase/sales conditions.
	ID []IdentifierType `json:"ID,omitempty"`
}

// A class to define a reference to an order line.
type OrderLineReferenceDetails struct {
	// An identifier for the referenced order line, assigned by the buyer.
	LineID []IdentifierType `json:"LineID"`
	// A code signifying the status of the referenced order line with respect to its original
	// state.
	LineStatusCode []CodeType `json:"LineStatusCode,omitempty"`
	// A reference to the Order containing the referenced order line.
	OrderReference []OrderReferenceDetails `json:"OrderReference,omitempty"`
	// An identifier for the referenced order line, assigned by the seller.
	SalesOrderLineID []IdentifierType `json:"SalesOrderLineID,omitempty"`
	// A universally unique identifier for this order line reference.
	UUID []IdentifierType `json:"UUID,omitempty"`
}

// A class to define a reference to an Order.
type OrderReferenceDetails struct {
	// Indicates whether the referenced Order is a copy (true) or the original (false).
	CopyIndicator []IndicatorType `json:"CopyIndicator,omitempty"`
	// Text used for tagging purchasing card transactions.
	CustomerReference []TextType `json:"CustomerReference,omitempty"`
	// A document associated with this reference to an Order.
	DocumentReference []DocumentReferenceDetails `json:"DocumentReference,omitempty"`
	// An identifier for this order reference, assigned by the buyer.
	ID []IdentifierType `json:"ID"`
	// The date on which the referenced Order was issued.
	IssueDate []DateType `json:"IssueDate,omitempty"`
	// The time at which the referenced Order was issued.
	IssueTime []TimeType `json:"IssueTime,omitempty"`
	// A code signifying the type of the referenced Order.
	OrderTypeCode []CodeType `json:"OrderTypeCode,omitempty"`
	// An identifier for this order reference, assigned by the seller.
	SalesOrderID []IdentifierType `json:"SalesOrderID,omitempty"`
	// A universally unique identifier for this order reference.
	UUID []IdentifierType `json:"UUID,omitempty"`
}

// A class describing identifiers or references relating to customs procedures.
type CustomsDeclarationDetails struct {
	// An identifier associated with customs related procedures.
	ID []IdentifierType `json:"ID"`
	// Describes the party issuing the customs declaration.
	IssuerParty []PartyDetails `json:"IssuerParty,omitempty"`
}

// A class to describe the condition or position of an object.
type StatusDetails struct {
	// Measurements that quantify the condition of the objects covered by the status.
	Condition []ConditionDetails `json:"Condition,omitempty"`
	// Specifies the status condition of the related object.
	ConditionCode []CodeType `json:"ConditionCode,omitempty"`
	// Text describing this status.
	Description []TextType `json:"Description,omitempty"`
	// Specifies an indicator relevant to a specific status.
	IndicationIndicator []IndicatorType `json:"IndicationIndicator,omitempty"`
	// A percentage meaningful in the context of this status.
	Percent []NumericType `json:"Percent,omitempty"`
	// The reference date for this status.
	ReferenceDate []DateType `json:"ReferenceDate,omitempty"`
	// The reference time for this status.
	ReferenceTime []TimeType `json:"ReferenceTime,omitempty"`
	// The reliability of this status, expressed as a percentage.
	ReliabilityPercent []NumericType `json:"ReliabilityPercent,omitempty"`
	// A sequence identifier for this status.
	SequenceID []IdentifierType `json:"SequenceID,omitempty"`
	// The reason for this status condition or position, expressed as text.
	StatusReason []TextType `json:"StatusReason,omitempty"`
	// The reason for this status condition or position, expressed as a code.
	StatusReasonCode []CodeType `json:"StatusReasonCode,omitempty"`
	// Provides any textual information related to this status.
	Text []TextType `json:"Text,omitempty"`
}

// A class to define a measurable condition of an object.
type ConditionDetails struct {
	// An identifier for the attribute that applies to the condition.
	AttributeID []IdentifierType `json:"AttributeID"`
	// Text describing the attribute that applies to the condition.
	Description []TextType `json:"Description,omitempty"`
	// The maximum value in a range of measurement for this condition.
	MaximumMeasure []MeasureType `json:"MaximumMeasure,omitempty"`
	// The measurement value.
	Measure []MeasureType `json:"Measure,omitempty"`
	// The minimum value in a range of measurement for this condition.
	MinimumMeasure []MeasureType `json:"MinimumMeasure,omitempty"`
}

// A class to describe a particular vehicle or vessel used for the conveyance of goods or
// persons.
type TransportMeansDetails struct {
	// An aircraft used for transport.
	AirTransport []AirTransportDetails `json:"AirTransport,omitempty"`
	// A code signifying the direction of this means of transport.
	DirectionCode []CodeType `json:"DirectionCode,omitempty"`
	// An identifier for the regular service schedule of this means of transport.
	JourneyID []IdentifierType `json:"JourneyID,omitempty"`
	// A vessel used for transport by water (not only by sea).
	MaritimeTransport []MaritimeTransportDetails `json:"MaritimeTransport,omitempty"`
	// A measurable dimension (length, mass, weight, or volume) of this means of transport.
	MeasurementDimension []DimensionDetails `json:"MeasurementDimension,omitempty"`
	// The party that owns this means of transport.
	OwnerParty []PartyDetails `json:"OwnerParty,omitempty"`
	// Equipment used for rail transport.
	RailTransport []RailTransportDetails `json:"RailTransport,omitempty"`
	// Text describing the country in which this means of transport is registered.
	RegistrationNationality []TextType `json:"RegistrationNationality,omitempty"`
	// An identifier for the country in which this means of transport is registered.
	RegistrationNationalityID []IdentifierType `json:"RegistrationNationalityID,omitempty"`
	// A vehicle used for road transport.
	RoadTransport []RoadTransportDetails `json:"RoadTransport,omitempty"`
	// The location within the means of transport where goods are to be or have been stowed.
	Stowage []StowageDetails `json:"Stowage,omitempty"`
	// A code signifying the service regularly provided by the carrier operating this means of
	// transport.
	TradeServiceCode []CodeType `json:"TradeServiceCode,omitempty"`
	// A code signifying the type of this means of transport (truck, vessel, etc.).
	TransportMeansTypeCode []CodeType `json:"TransportMeansTypeCode,omitempty"`
}

// A class to identify a specific aircraft used for transportation.
type AirTransportDetails struct {
	// An identifer for a specific aircraft.
	AircraftID []IdentifierType `json:"AircraftID"`
}

// A class to describe a vessel used for transport by water (including sea, river, and
// canal).
type MaritimeTransportDetails struct {
	// Gross tonnage is calculated by measuring a ship's volume (from keel to funnel, to the
	// outside of the hull framing) and applying a mathematical formula and is used to determine
	// things such as a ship's manning regulations, safety rules, registration fees and port
	// dues.
	GrossTonnageMeasure []MeasureType `json:"GrossTonnageMeasure,omitempty"`
	// Net tonnage is calculated by measuring a ship's internal volume and applying a
	// mathematical formula and is used to calculate the port duties.
	NetTonnageMeasure []MeasureType `json:"NetTonnageMeasure,omitempty"`
	// The radio call sign of the vessel.
	RadioCallSignID []IdentifierType `json:"RadioCallSignID,omitempty"`
	// The certificate issued to the ship by the ships registry in a given flag state.
	RegistryCertificateDocumentReference []DocumentReferenceDetails `json:"RegistryCertificateDocumentReference,omitempty"`
	// The port in which a vessel is registered or permanently based.
	RegistryPortLocation []LocationDetails `json:"RegistryPortLocation,omitempty"`
	// Information about what services a vessel will require when it arrives at a port, such as
	// refueling, maintenance, waste disposal etc.
	ShipsRequirements []TextType `json:"ShipsRequirements,omitempty"`
	// An identifier for a specific vessel.
	VesselID []IdentifierType `json:"VesselID,omitempty"`
	// The name of the vessel.
	VesselName []TextType `json:"VesselName,omitempty"`
}

// A class defining details about a train wagon used as a means of transport.
type RailTransportDetails struct {
	// An identifier for the rail car on the train used as the means of transport.
	RailCarID []IdentifierType `json:"RailCarID,omitempty"`
	// An identifier for the train used as the means of transport.
	TrainID []IdentifierType `json:"TrainID"`
}

// A class for identifying a vehicle used for road transport.
type RoadTransportDetails struct {
	// The license plate identifier of this vehicle.
	LicensePlateID []IdentifierType `json:"LicensePlateID"`
}

// A class to describe a location on board a means of transport where specified goods or
// transport equipment have been stowed or are to be stowed.
type StowageDetails struct {
	// Text describing the location.
	Location []TextType `json:"Location,omitempty"`
	// An identifier for the location.
	LocationID []IdentifierType `json:"LocationID,omitempty"`
	// A measurable dimension (length, mass, weight, or volume) of this stowage.
	MeasurementDimension []DimensionDetails `json:"MeasurementDimension,omitempty"`
}

// A class to describe a delivery unit.
type DeliveryUnitDetails struct {
	// The quantity of ordered Items that constitutes a batch for delivery purposes.
	BatchQuantity []QuantityType `json:"BatchQuantity"`
	// The quantity of units in the Delivery Unit expressed in the units used by the consumer.
	ConsumerUnitQuantity []QuantityType `json:"ConsumerUnitQuantity,omitempty"`
	// An indication that the transported goods are subject to an international regulation
	// concerning the carriage of dangerous goods (true) or not (false).
	HazardousRiskIndicator []IndicatorType `json:"HazardousRiskIndicator,omitempty"`
}

// A class to define the price of an item as a percentage of the price of a different item.
type DependentPriceReferenceDetails struct {
	// A reference to a line that the price is depended of.
	DependentLineReference []LineReferenceDetails `json:"DependentLineReference,omitempty"`
	// The reference location for this dependent price reference.
	LocationAddress []AddressDetails `json:"LocationAddress,omitempty"`
	// The percentage by which the price of the different item is multiplied to calculate the
	// price of the item.
	Percent []NumericType `json:"Percent,omitempty"`
}

// A class to describe a price list.
type PriceListDetails struct {
	// An identifier for this price list.
	ID []IdentifierType `json:"ID,omitempty"`
	// The previous price list.
	PreviousPriceList []PriceListDetails `json:"PreviousPriceList,omitempty"`
	// A code signifying whether this price list is an original, copy, revision, or cancellation.
	StatusCode []CodeType `json:"StatusCode,omitempty"`
	// A period during which this price list is valid.
	ValidityPeriod []PeriodDetails `json:"ValidityPeriod,omitempty"`
}

// A class for describing the terms and conditions applying to the delivery of goods.
type DeliveryTermsDetails struct {
	// An allowance or charge covered by these delivery terms.
	AllowanceCharge []AllowanceChargeDetails `json:"AllowanceCharge,omitempty"`
	// The monetary amount covered by these delivery terms.
	Amount []AmountType `json:"Amount,omitempty"`
	// The location for the contracted delivery.
	DeliveryLocation []LocationDetails `json:"DeliveryLocation,omitempty"`
	// An identifier for this description of delivery terms.
	ID []IdentifierType `json:"ID,omitempty"`
	// A description of responsibility for risk of loss in execution of the delivery, expressed
	// as text.
	LossRisk []TextType `json:"LossRisk,omitempty"`
	// A code that identifies one of various responsibilities for loss risk in the execution of
	// the delivery.
	LossRiskResponsibilityCode []CodeType `json:"LossRiskResponsibilityCode,omitempty"`
	// A description of any terms or conditions relating to the delivery items.
	SpecialTerms []TextType `json:"SpecialTerms,omitempty"`
}

// A class to describe a price extension, calculated by multiplying the price per unit by
// the quantity of items.
type PriceExtensionDetails struct {
	// The amount of this price extension.
	Amount []AmountType `json:"Amount"`
	// A total amount of taxes of a particular kind applicable to this price extension.
	TaxTotal []TaxTotalDetails `json:"TaxTotal,omitempty"`
}

// A class to describe the despatching of goods (their pickup for delivery).
type DespatchDetails struct {
	// The actual despatch (pickup) date.
	ActualDespatchDate []DateType `json:"ActualDespatchDate,omitempty"`
	// The actual despatch (pickup) time.
	ActualDespatchTime []TimeType `json:"ActualDespatchTime,omitempty"`
	// The party carrying the goods.
	CarrierParty []PartyDetails `json:"CarrierParty,omitempty"`
	// The primary contact for this despatch (pickup).
	Contact []ContactDetails `json:"Contact,omitempty"`
	// The address of the despatch (pickup).
	DespatchAddress []AddressDetails `json:"DespatchAddress,omitempty"`
	// The location of the despatch (pickup).
	DespatchLocation []LocationDetails `json:"DespatchLocation,omitempty"`
	// The party despatching the goods.
	DespatchParty []PartyDetails `json:"DespatchParty,omitempty"`
	// The estimated despatch (pickup) date.
	EstimatedDespatchDate []DateType `json:"EstimatedDespatchDate,omitempty"`
	// The period estimated for the despatch (pickup) of goods.
	EstimatedDespatchPeriod []PeriodDetails `json:"EstimatedDespatchPeriod,omitempty"`
	// The estimated despatch (pickup) time.
	EstimatedDespatchTime []TimeType `json:"EstimatedDespatchTime,omitempty"`
	// The date guaranteed for the despatch (pickup).
	GuaranteedDespatchDate []DateType `json:"GuaranteedDespatchDate,omitempty"`
	// The time guaranteed for the despatch (pickup).
	GuaranteedDespatchTime []TimeType `json:"GuaranteedDespatchTime,omitempty"`
	// An identifier for this despatch event.
	ID []IdentifierType `json:"ID,omitempty"`
	// Text describing any special instructions applying to the despatch (pickup).
	Instructions []TextType `json:"Instructions,omitempty"`
	// A party to be notified of this despatch (pickup).
	NotifyParty []PartyDetails `json:"NotifyParty,omitempty"`
	// An identifier for the release of the despatch used as security control or cargo control
	// (pick-up).
	ReleaseID []IdentifierType `json:"ReleaseID,omitempty"`
	// The despatch (pickup) date requested, normally by the buyer.
	RequestedDespatchDate []DateType `json:"RequestedDespatchDate,omitempty"`
	// The period requested for the despatch (pickup) of goods.
	RequestedDespatchPeriod []PeriodDetails `json:"RequestedDespatchPeriod,omitempty"`
	// The despatch (pickup) time requested, normally by the buyer.
	RequestedDespatchTime []TimeType `json:"RequestedDespatchTime,omitempty"`
}

// A class to describe a pickup for delivery.
type PickupDetails struct {
	// The actual pickup date.
	ActualPickupDate []DateType `json:"ActualPickupDate,omitempty"`
	// The actual pickup time.
	ActualPickupTime []TimeType `json:"ActualPickupTime,omitempty"`
	// The earliest pickup date.
	EarliestPickupDate []DateType `json:"EarliestPickupDate,omitempty"`
	// The earliest pickup time.
	EarliestPickupTime []TimeType `json:"EarliestPickupTime,omitempty"`
	// An identifier for this pickup.
	ID []IdentifierType `json:"ID,omitempty"`
	// The latest pickup date.
	LatestPickupDate []DateType `json:"LatestPickupDate,omitempty"`
	// The latest pickup time.
	LatestPickupTime []TimeType `json:"LatestPickupTime,omitempty"`
	// The pickup location.
	PickupLocation []LocationDetails `json:"PickupLocation,omitempty"`
	// The party responsible for picking up a delivery.
	PickupParty []PartyDetails `json:"PickupParty,omitempty"`
}

// A class for describing the terms of a trade agreement.
type TradingTermsDetails struct {
	// The address at which these trading terms apply.
	ApplicableAddress []AddressDetails `json:"ApplicableAddress,omitempty"`
	// Text describing the terms of a trade agreement.
	Information []TextType `json:"Information,omitempty"`
	// A reference quoting the basis of the terms
	Reference []TextType `json:"Reference,omitempty"`
}

// A class to describe a device (a transport equipment seal) for securing the doors of a
// shipping container.
type TransportEquipmentSealDetails struct {
	// The condition of this transport equipment seal.
	Condition []TextType `json:"Condition,omitempty"`
	// An identifier for this transport equipment seal.
	ID []IdentifierType `json:"ID"`
	// The role of the sealing party.
	SealingPartyType []TextType `json:"SealingPartyType,omitempty"`
	// A code signifying the type of party that issues and is responsible for this transport
	// equipment seal.
	SealIssuerTypeCode []CodeType `json:"SealIssuerTypeCode,omitempty"`
	// A code signifying the condition of this transport equipment seal.
	SealStatusCode []CodeType `json:"SealStatusCode,omitempty"`
}

// A class to describe an environmental emission.
type EnvironmentalEmissionDetails struct {
	// Text describing this environmental emission.
	Description []TextType `json:"Description,omitempty"`
	// A method used to calculate the amount of this emission.
	EmissionCalculationMethod []EmissionCalculationMethodDetails `json:"EmissionCalculationMethod,omitempty"`
	// A code specifying the type of environmental emission.
	EnvironmentalEmissionTypeCode []CodeType `json:"EnvironmentalEmissionTypeCode"`
	// A value measurement for the environmental emission.
	ValueMeasure []MeasureType `json:"ValueMeasure"`
}

// A class to define how an environmental emission is calculated.
type EmissionCalculationMethodDetails struct {
	// A code signifying the method used to calculate the emission.
	CalculationMethodCode []CodeType `json:"CalculationMethodCode,omitempty"`
	// A code signifying whether a piece of transport equipment is full, partially full, or
	// empty. This indication is used as a parameter when calculating the environmental emission.
	FullnessIndicationCode []CodeType `json:"FullnessIndicationCode,omitempty"`
	// A start location from which an environmental emission is calculated.
	MeasurementFromLocation []LocationDetails `json:"MeasurementFromLocation,omitempty"`
	// An end location to which an environmental emission is calculated.
	MeasurementToLocation []LocationDetails `json:"MeasurementToLocation,omitempty"`
}

// A class to specify which day of the week a transport service is operational.
type ServiceFrequencyDetails struct {
	// A day of the week, expressed as code.
	WeekDayCode []CodeType `json:"WeekDayCode"`
}

// A class to describe an application-level response to a document.
type ResponseDetails struct {
	// Text describing this response.
	Description []TextType `json:"Description,omitempty"`
	// The date upon which this response is valid.
	EffectiveDate []DateType `json:"EffectiveDate,omitempty"`
	// The time at which this response is valid.
	EffectiveTime []TimeType `json:"EffectiveTime,omitempty"`
	// An identifier for the section (or line) of the document to which this response applies.
	ReferenceID []IdentifierType `json:"ReferenceID,omitempty"`
	// A code signifying the type of response.
	ResponseCode []CodeType `json:"ResponseCode,omitempty"`
	// A status report associated with this response.
	Status []StatusDetails `json:"Status,omitempty"`
}

// A class to define a monetary total.
type MonetaryTotalDetails struct {
	// The total monetary amount of all allowances.
	AllowanceTotalAmount []AmountType `json:"AllowanceTotalAmount,omitempty"`
	// The total monetary amount of all charges.
	ChargeTotalAmount []AmountType `json:"ChargeTotalAmount,omitempty"`
	// The monetary amount of an extended transaction line, net of tax and settlement discounts,
	// but inclusive of any applicable rounding amount.
	LineExtensionAmount []AmountType `json:"LineExtensionAmount,omitempty"`
	// The amount of the monetary total to be paid, expressed in an alternative currency.
	PayableAlternativeAmount []AmountType `json:"PayableAlternativeAmount,omitempty"`
	// The amount of the monetary total to be paid.
	PayableAmount []AmountType `json:"PayableAmount"`
	// The rounding amount (positive or negative) added to produce the line extension amount.
	PayableRoundingAmount []AmountType `json:"PayableRoundingAmount,omitempty"`
	// The total prepaid monetary amount.
	PrepaidAmount []AmountType `json:"PrepaidAmount,omitempty"`
	// The monetary amount of an extended transaction line, exclusive of taxes.
	TaxExclusiveAmount []AmountType `json:"TaxExclusiveAmount,omitempty"`
	// The monetary amount including taxes; the sum of payable amount and prepaid amount.
	TaxInclusiveAmount []AmountType `json:"TaxInclusiveAmount,omitempty"`
}

// A class to describe a payment.
type PaymentDetails struct {
	// An identifier for this payment.
	ID []IdentifierType `json:"ID,omitempty"`
	// An identifier for the payment instruction.
	InstructionID []IdentifierType `json:"InstructionID,omitempty"`
	// The amount of this payment.
	PaidAmount []AmountType `json:"PaidAmount,omitempty"`
	// The date on which this payment was made.
	PaidDate []DateType `json:"PaidDate,omitempty"`
	// The time at which this payment was made.
	PaidTime []TimeType `json:"PaidTime,omitempty"`
	// The date on which this payment was received.
	ReceivedDate []DateType `json:"ReceivedDate,omitempty"`
}

// A class to define a reference to a procurement project.
type ProjectReferenceDetails struct {
	// An identifier for the referenced project.
	ID []IdentifierType `json:"ID"`
	// The date on which the referenced project was issued.
	IssueDate []DateType `json:"IssueDate,omitempty"`
	// A universally unique identifier for the referenced project.
	UUID []IdentifierType `json:"UUID,omitempty"`
	// A specific phase of work in the referenced project.
	WorkPhaseReference []WorkPhaseReferenceDetails `json:"WorkPhaseReference,omitempty"`
}

// A class that refers to a phase of work. Used for instance to specify what part of the
// contract the billing is referring to.
type WorkPhaseReferenceDetails struct {
	// The date on which this phase of work ends.
	EndDate []DateType `json:"EndDate,omitempty"`
	// An identifier for this phase of work.
	ID []IdentifierType `json:"ID,omitempty"`
	// The progress percentage of the work phase.
	ProgressPercent []NumericType `json:"ProgressPercent,omitempty"`
	// The date on which this phase of work begins.
	StartDate []DateType `json:"StartDate,omitempty"`
	// A reference to a document regarding the work order for the project in which this phase of
	// work takes place.
	WorkOrderDocumentReference []DocumentReferenceDetails `json:"WorkOrderDocumentReference,omitempty"`
	// Text describing this phase of work.
	WorkPhase []TextType `json:"WorkPhase,omitempty"`
	// A code signifying this phase of work.
	WorkPhaseCode []CodeType `json:"WorkPhaseCode,omitempty"`
}

// A container for all extensions present in the document.
type UBLExtensions struct {
	// A single extension for private use.
	UBLExtension []UBLExtension `json:"UBLExtension"`
}

// A single extension for private use.
type UBLExtension struct {
	// An agency that maintains one or more Extensions.
	ExtensionAgencyID []IdentifierType `json:"ExtensionAgencyID,omitempty"`
	// The name of the agency that maintains the Extension.
	ExtensionAgencyName []TextType `json:"ExtensionAgencyName,omitempty"`
	// A URI for the Agency that maintains the Extension.
	ExtensionAgencyURI []IdentifierType `json:"ExtensionAgencyURI,omitempty"`
	// The definition of the extension content.
	ExtensionContent []map[string]interface{} `json:"ExtensionContent"`
	// A description of the reason for the Extension.
	ExtensionReason []TextType `json:"ExtensionReason,omitempty"`
	// A code for reason the Extension is being included.
	ExtensionReasonCode []CodeType `json:"ExtensionReasonCode,omitempty"`
	// A URI for the Extension.
	ExtensionURI []IdentifierType `json:"ExtensionURI,omitempty"`
	// The version of the Extension.
	ExtensionVersionID []IdentifierType `json:"ExtensionVersionID,omitempty"`
	// An identifier for the Extension assigned by the creator of the extension.
	ID []IdentifierType `json:"ID,omitempty"`
	// A name for the Extension assigned by the creator of the extension.
	Name []TextType `json:"Name,omitempty"`
}
