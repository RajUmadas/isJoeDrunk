"""`main` is the top level module for your Flask application."""

from time import sleep
from flask import Flask, render_template, redirect
import os
from pymongo import MongoClient

APP = Flask(__name__)
client = MongoClient(os.environ['DB_PORT_27017_TCP_ADDR'], 27017)
db = client.db

@APP.route('/')
def show_answer():
    """Main entry point for the answer"""
    if db.answers.find().count()==0:
        ans = "asdfasfasdf"
    else:
        ans = db.answers.find().sort([("_id", -1)]).limit(1)[0]['answer']

    return render_template('page.html', answer=ans)

@APP.route('/set')
def easyset():
    """Buttons for when joe is drunk"""
    return render_template('set.html')

@APP.route('/set/no')
def set_no():
    """Setting No"""
    db.answers.insert_one({'answer': 'No'})
    sleep(0.5)
    return redirect('/')

@APP.route('/set/yes')
def set_yes():
    """Setting Yes"""
    db.answers.insert_one({'answer': 'Yes'})
    sleep(0.5)
    return redirect('/')


if __name__ == "__main__":
    APP.run(host='0.0.0.0', port=80,  debug=False)
