"""`main` is the top level module for your Flask application."""

from time import sleep
from flask import Flask, render_template, redirect

from models.answers import Answers
from models.detained import Detained
APP = Flask(__name__)

@APP.route('/')
def show_answer():
    """Main entry point for the answer"""
    recent_answer = Answers().query().order(-Answers.date).fetch(1)
    ans = None
    if len(recent_answer) != 0:
        ans = recent_answer[0].answer
    return render_template('page.html', answer=ans)

@APP.route('/set')
def easyset():
    """Buttons for when joe is drunk"""
    return render_template('set.html')

@APP.route('/set/no')
def set_no():
    """Setting No"""
    Answers(answer='No').put()
    sleep(0.5)
    return redirect('/')

@APP.route('/set/yes')
def set_yes():
    """Setting Yes"""
    Answers(answer='Yes').put()
    sleep(0.5)
    return redirect('/')

@APP.route('/set/det/no')
def set_det_no():
    return redirect('/')
    """Setting No"""
    Detained(answer='No').put()
    sleep(0.5)
    return redirect('/')

@APP.route('/set/det/yes')
def set_det_yes():
    """Setting Yes"""
    return redirect('/')
    Detained(answer='Yes').put()
    sleep(0.5)
    return redirect('/')

@APP.route('/admin/del')
def del_all():
    """Clearing Google Datastore"""
    for i in Answers().query():
        i.key.delete()
    for i in Detained().query():
        i.key.delete()
    return redirect('/')

@APP.errorhandler(404)
def page_not_found(dummy):
    """Return a custom 404 error."""
    return render_template('error.html', reason="Aint Here"), 404

@APP.errorhandler(500)
def application_error(dummy):
    """Return a custom 500 error."""
    return render_template('error.html', reason="Whoaaaa"), 500
