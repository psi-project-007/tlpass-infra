# tlpass-infra

> **Pre-alpha: do not store real passwords.** Nothing here has been audited yet, and the infrastructure can still change.

The TLPass infrastructure on AWS, written with the AWS CDK in Go. It defines the dev and prod stacks that run the API: Cognito, DynamoDB, API Gateway, Lambda and their alarms, all in `eu-central-1`.

GitHub Actions deploys it with short-lived OIDC credentials, so no AWS access key is stored anywhere. Dev deploys from `main`; prod deploys only from a release tag after the owner approves.

## How it fits

| Repository | Role |
| --- | --- |
| tlpass-crypto | The cryptography, published as `@psi-project-007/tlpass-crypto` |
| tlpass-web | The web portal |
| tlpass-api | The API (Go Lambda handlers), which stores only encrypted data |
| **tlpass-infra** (this one) | The AWS infrastructure (CDK in Go) |

## Licence

[AGPL-3.0](LICENSE).

## Security

Please don't report vulnerabilities in public issues. Follow the [security policy](https://github.com/psi-project-007/.github/blob/main/SECURITY.md).
