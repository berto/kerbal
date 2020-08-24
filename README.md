# Kerbal.me

> Kerbal avatar generator

Kerbal.me is a fan-made website, it is in no way affiliated with the KSP team.

## Artwork

The artwork is stored in a AWS S3 account for the app under the following folders:

```
  /images
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

The application uses Amazon Services:

- Client: hosted on AWS S3 through AWS Cloudfront
- Server: AWS Lambda available through AWS API Gateway
- Images: AWS S3

The server connects to AWS S3 to load the images using the following envs:

```
AWS_BUCKET=XXXXXXXXXX **defaults to kerbal.me**
AWS_REGION=XXXXXXXXXX **defaults to us-west-2**
```

### Docker

The website can be run as a standalone application under the `docker` branch. It requires AWS credentials to be set up. If `AWS_ACCESS_KEY_ID` or `AWS_SECRET_ACCESS_KEY` are not provided, it looks for shared credentials under the `AWS_PROFILE` name in `~/.aws/credentials`

```
AWS_PROFILE=XXXXXXXXXX **defaults to kerbal.me**
```

To run with docker:

`git checkout docker`

`docker build -t kerbal .`

`docker run -p 3000:3000 -e AWS_ACCESS_KEY_ID=access_key_id -e AWS_SECRET_ACCESS_KEY=secret_access_key kerbal`
