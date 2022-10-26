import dash
from dash import Dash, html
import dash_bootstrap_components as dbc
from flask_simplelogin import login_required

from src.flaskr.blueprints.dash.component.header import header


def init_app(app):
    dash_app = Dash(
        __name__,
        server=app,
        url_base_pathname='/dash/',
        external_stylesheets=[dbc.themes.BOOTSTRAP, dbc.icons.FONT_AWESOME],
        use_pages=True
    )

    dash_app.layout = html.Div(children=[
        header(),
        dbc.Container(children=[
            dash.page_container
        ]),
    ])

    dash_app.index = login_required(dash_app.index)

    @app.route("/dash")
    def my_dash_app():
        return dash_app.index()
