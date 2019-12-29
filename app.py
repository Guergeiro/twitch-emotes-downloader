from bs4 import BeautifulSoup
import requests
from zipfile import ZipFile
import argparse
from os import remove as fileRemove

parser = argparse.ArgumentParser()

parser.add_argument(
    "-U",
    "--url",
    help=
    "URL of the emotes you want.\nDefault is https://www.twitchmetrics.net/emotes",
    default="https://www.twitchmetrics.net/emotes")

parser.add_argument("-O",
                    "--output",
                    help="Output filename of zip.\nDefault is 'output.zip'",
                    default="output.zip")

args = parser.parse_args()


def requestContent(url):
    return requests.get(url).content


def writeZip(file, emotes):
    # Opens a zip to write
    with ZipFile(f"{file}", "w") as zipFile:
        # Iterates over all images
        for emote in emotes:
            # Opens a {imageName}.png to write image
            with open(f"{emote.get('name')}.png", "wb") as imageFile:
                imageFile.write(emote.get("content"))
            imageFile.close()
            zipFile.write(f"{emote.get('name')}.png")
            fileRemove(f"{emote.get('name')}.png")
    zipFile.close()


def parsePage(page):
    parsed_page = BeautifulSoup(page, "html.parser")
    return parsed_page


page = requests.get(args.url)

parsed_html = BeautifulSoup(page.content, "html.parser")

names = parsed_html.body.find_all("samp")

emotes = []

for name in names:
    emotes.append({
        "name":
        name.text,
        "content":
        requestContent(name.find_previous_sibling("div").find("img")["src"])
    })

writeZip(args.output, emotes)