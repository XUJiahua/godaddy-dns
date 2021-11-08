/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package godaddy

type DomainDnssec struct {
	// This identifies the cryptographic algorithm used to generate the signature<br/><ul><li><strong style='margin-left: 12px;'>RSAMD5</strong> - [01] DRSA/MD5 </li><li><strong style='margin-left: 12px;'>DSA</strong> - [03] DSA/SHA1</li><li><strong style='margin-left: 12px;'>RSASHA1</strong> - [05] RSA/SHA-1</li><li><strong style='margin-left: 12px;'>DSA_NSEC3_SHA1</strong> - [06] DSA-NSEC3-SHA1</li><li><strong style='margin-left: 12px;'>RSASHA1_NSEC3_SHA1</strong> - [07] RSASHA1-NSEC3-SHA1</li><li><strong style='margin-left: 12px;'>RSASHA256</strong> - [08] RSA/SHA-256</li><li><strong style='margin-left: 12px;'>RSASHA512</strong> - [10] RSA/SHA-512</li><li><strong style='margin-left: 12px;'>ECC_GOST</strong> - [12] GOST R 34.10-2001</li><li><strong style='margin-left: 12px;'>ECDSAP256SHA256</strong> - [13] ECDSA Curve P-256 with SHA-256</li><li><strong style='margin-left: 12px;'>ECDSAP384SHA384</strong> - [14] ECDSA Curve P-384 with SHA-384</li><li><strong style='margin-left: 12px;'>ED25519</strong> - [15] Ed25519</li><li><strong style='margin-left: 12px;'>ED448</strong> - [16] Ed448</li></ul>
	Algorithm string `json:"algorithm"`
	// This is an integer value less than 65536 used to identify the DNSSEC record for the domain name.
	KeyTag int32 `json:"keyTag,omitempty"`
	// This identifies the algorithm used to construct the digest<br/><ul><li><strong style='margin-left: 12px;'>SHA1</strong> - [01] SHA-1</li><li><strong style='margin-left: 12px;'>SHA256</strong> - [02] SHA-256</li><li><strong style='margin-left: 12px;'>GOST</strong> - [03] GOST R 34.11-94</li><li><strong style='margin-left: 12px;'>SHA384</strong> - [04] SHA-384</li></ul>
	DigestType string `json:"digestType,omitempty"`
	// The digest is an alpha-numeric value
	Digest string `json:"digest,omitempty"`
	// This identifies the key type; either a Zone-Signing Key or a Key-Signing Key<br/><ul><li><strong style='margin-left: 12px;'>ZSK</strong> - [256] Zone-Signing Key</li><li><strong style='margin-left: 12px;'>KSK</strong> - [257] Key-Signing Key</li></ul>
	Flags string `json:"flags,omitempty"`
	// Registries use this value to encrypt DS records. Decryption requires a matching public key
	PublicKey string `json:"publicKey,omitempty"`
	// This specifies the validity period for the signature. The value is expressed in seconds. You can use any integer value larger than zero
	MaxSignatureLife int32 `json:"maxSignatureLife,omitempty"`
}