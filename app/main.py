"""`main` is the top level module for your Flask application."""

from time import sleep
from flask import Flask, render_template, redirect

from models.answers import Answers
from models.detained import Detained
APP = Flask(__name__)

@APP.route('/')
def show_answer():
    """Main entry point for the answer"""
    ans = None
    return render_template('page.html', answer=ans)

@APP.route('/set')
def easyset():
    """Buttons for when joe is drunk"""
    return render_template('set.html')


if __name__ == "__main__":
    APP.run(host='0.0.0.0', port=80,  debug=False)
