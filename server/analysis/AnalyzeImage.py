import os
import sys
import requests
import time
import matplotlib.pyplot as plt
from matplotlib.patches import Polygon
from io import BytesIO

import Env

# Add your Computer Vision subscription key and endpoint to your environment variables.
if Env.COMPUTER_VISION_SUBSCRIPTION_KEY != "":
    subscription_key = Env.COMPUTER_VISION_SUBSCRIPTION_KEY
else:
    print("\nSet the COMPUTER_VISION_SUBSCRIPTION_KEY environment variable.\n**Restart your shell or IDE for changes to take effect.**")
    sys.exit()

if Env.COMPUTER_VISION_ENDPOINT != "":
    endpoint = Env.COMPUTER_VISION_ENDPOINT

text_recognition_url = endpoint + "vision/v2.0/read/core/asyncBatchAnalyze"

# Set imageURL to the URL of an image that you want to analyze.
imageURL = "https://upload.wikimedia.org/wikipedia/commons/d/dd/Cursive_Writing_on_Notebook_paper.jpg"
if len(sys.argv) == 2:
    imageURL = sys.argv[1]

headers = {'Ocp-Apim-Subscription-Key': subscription_key}
data = {'url': imageURL}
response = requests.post(
    text_recognition_url, headers=headers, json=data)
response.raise_for_status()

# Extracting text requires two API calls: One call to submit the
# image for processing, the other to retrieve the text found in the image.

# Holds the URI used to retrieve the recognized text.
operation_url = response.headers["Operation-Location"]

# The recognized text isn't immediately available, so poll to wait for completion.
analysis = {}
poll = True
while (poll):
    response_final = requests.get(
        response.headers["Operation-Location"], headers=headers)
    analysis = response_final.json()
    # print(analysis)
    time.sleep(1)
    if ("recognitionResults" in analysis):
        poll = False
    if ("status" in analysis and analysis['status'] == 'Failed'):
        poll = False

# print(sys.argv)
print(analysis)
