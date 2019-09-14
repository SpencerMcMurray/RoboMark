import os
import sys
import requests
import time
import matplotlib.pyplot as plt
import pprint
from matplotlib.patches import Polygon
from io import BytesIO

import Env


# from PIL import Image

# Add your Computer Vision subscription key and endpoint to your environment variables.
if Env.COMPUTER_VISION_SUBSCRIPTION_KEY != "":
    subscription_key = Env.COMPUTER_VISION_SUBSCRIPTION_KEY
else:
    print("\nSet the COMPUTER_VISION_SUBSCRIPTION_KEY environment variable.\n**Restart your shell or IDE for changes to take effect.**")
    sys.exit()

if Env.COMPUTER_VISION_ENDPOINT != "":
    endpoint = Env.COMPUTER_VISION_ENDPOINT

text_recognition_url = endpoint + "vision/v2.0/read/core/asyncBatchAnalyze"

# Set image_url to the URL of an image that you want to analyze.
image_url = "https://upload.wikimedia.org/wikipedia/commons/d/dd/Cursive_Writing_on_Notebook_paper.jpg"

headers = {'Ocp-Apim-Subscription-Key': subscription_key}
data = {'url': image_url}
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

polygons = []
if ("recognitionResults" in analysis):
    # Extract the recognized text, with bounding boxes.
    polygons = [(line["boundingBox"], line["text"])
                for line in analysis["recognitionResults"][0]["lines"]]

# Display the image and overlay it with the extracted text.

# plt.figure(figsize=(15, 15))
# image = Image.open(BytesIO(requests.get(image_url).content))
# ax = plt.imshow(image)
# for polygon in polygons:
#     vertices = [(polygon[0][i], polygon[0][i+1])
#                 for i in range(0, len(polygon[0]), 2)]
#     text = polygon[1]
#     patch = Polygon(vertices, closed=True, fill=False, linewidth=2, color='y')
#     ax.axes.add_patch(patch)
#     plt.text(vertices[0][0], vertices[0][1], text, fontsize=20, va="top")

pprint.pprint(analysis, indent=1)
