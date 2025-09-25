# Using Build Secrets with Okteto

This sample demonstrates how to securely use **build secrets** in your Okteto development environments.  

A **build secret** is any sensitive piece of information (e.g., passwords, API tokens) that your application needs during its build process.  

⚠️ **Important:** Do not pass secrets using build arguments or environment variables. These methods persist data in the final image, which is insecure. Instead, use **secret mounts** , which expose secrets only during the build process.

---

## Typical Use Cases

Common scenarios where build secrets are needed include:

- Tokens for private registries (npm, Artifactory, etc.)
- API keys
- Passwords for internal services

---

## How This Sample Works

This example uses the variable **`$NAME`** to illustrate how to configure build secrets in a simple way.  

`$NAME` can be defined in multiple ways:

- As a **local environment variable**
- As a **CLI parameter**
- As an [**Okteto Variable**](https://www.okteto.com/docs/core/okteto-variables/)

✅ **Recommendation:** Use **Okteto Variables**. They ensure a consistent experience across your team and reduce manual setup steps.


1. Create an Okteto Variable named `NAME` in your Okteto instance and give it a value (e.g. `cindy`)
2. Deploy your environment `okteto deploy --wait`
3. Once the environment is up, Okteto will expose an endpoint, retrive it by running `okteto endpoint` or by checking your Okteto Dashboard 
4. You can then open that endpoint in your browser to verify that it's displaying your secret name.

