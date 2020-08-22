# Kerbal.me

> Kerbal avatar generator

Kerbal.me is a fan-made website, it is in no way affiliated with the KSP team.

## Artwork

The artwork is stored in a AWS S3 account for the app under the following folders:

```
  /suit
  /color
  /eyes
  /mouth
  /hair
  /facial-hair
  /glasses
  /extras
  /kerbals
```

## Contributing

[Read Contributing Guideline](./contributing.md)

## Local Development

The application connects to AWS S3 to load the images using the following envs:

```
AWS_PROFILE=XXXXXXXXXX **defaults to kerbal.me**
AWS_BUCKET=XXXXXXXXXX **defaults to kerbal.me**
AWS_REGION=XXXXXXXXXX **defaults to us-west-2**
AWS_ACCESS_KEY_ID=XXXXXXXXXX
AWS_SECRET_ACCESS_KEY=XXXXXXXXXX
```

If `AWS_ACCESS_KEY_ID` or `AWS_SECRET_ACCESS_KEY` are not provided, it looks for shared credentials under the `AWS_PROFILE` name in `~/.aws/credentials`

To run with docker:

`docker build -t kerbal .`

`docker run -p 3000:3000 -e AWS_ACCESS_KEY_ID=access_key_id -e AWS_SECRET_ACCESS_KEY=secret_access_key kerbal`
