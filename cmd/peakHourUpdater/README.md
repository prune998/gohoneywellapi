# PeakHourUpdater

## Description

This app is a CLI that acts as a git client that:

1. ask for the new schedule to add
1. clone or pull the latest version of the git repo
1. update the file, push, PR
1. allow you to review the change and merge the PR

then the json schedule file will be available in Github Pages

## Usage

./peakHourUpdater --startDate <> --endDate <> [--file <> --repo <>]