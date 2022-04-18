# pgpkeygen

This quick script should be ran as a part of new user pipeline.

When a new user gets registered in the system, the username and email are pulled from the oidc provider. These are simply processed as inputs (coming from env vars) to this pgp key generator, and a templated global gitconfig is prepopluated with these values.

As the program does not create a keyring, don't forge to `gpg --import /path/to/privatekey` during the next step, and delete afterwards.
