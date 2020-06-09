#!/usr/bin/env bash

mkdir -p ~/.terraform.d/plugins
cp $(which terraform-provider-netlify) ~/.terraform.d/plugins/terraform-provider-netlify_v0.5.0

