from dash import html
import dash

dash.register_page(__name__)


def layout():
    return html.H1("This is our custom 404 content")
