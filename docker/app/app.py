"""Simple web application that takes JSON {word: "STRING", "count": NUM}
as a parameter and returns string with that "word" and emoji symbol that it's
mean NUM times
"""


from flask import Flask, render_template, request
from random import randint
import emoji

app = Flask(__name__)


@app.route("/<emoji_name>")
def preview(emoji_name):
    string = emoji.emojize(f":{emoji_name}:")
    return render_template("preview.html", content=string)


@app.route("/", methods=["POST", "GET"])
def emojis():
    """Generate a decorated string based on a JSON object, or return a static
    page, depending on the method header in the request.
    """
    emo_arr = ["elephant", "turtle", "mouse", "cat", "dog", "panda", "crocodile"]
    if request.method == "POST":
        content = request.get_json(force=True)
        emo_nm = emo_arr[randint(0, len(emo_arr) - 1)]
        word = content['word']
        count = content['count']
        string = ""
        for _ in range(int(count)):
            string += f":{emo_nm}:{word} \n"
        return emoji.emojize(string + '\n')
    return render_template("index.html")

@app.route("/description")
def description():
	return render_template('description.html')


if __name__ == "__main__":
    app.run(debug=True)
