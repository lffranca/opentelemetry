import flask
import dash
from dash import html, dcc, callback, Input, Output

dash.register_page(__name__, order=1, path_template="/sales/<client_id>")


def layout(client_id=None):
    # print(flask.request.headers)

    return html.Div(children=[
        html.H1(children='This is our sales page'),
        html.Div([
            "Select a city: ",
            dcc.RadioItems(['New York City', 'Montreal','San Francisco'],
                           'Montreal',
                           id='sales-input')
        ]),
        html.Br(),
        html.Div(id='sales-output'),
    ])


@callback(
    Output(component_id='sales-output', component_property='children'),
    Input(component_id='sales-input', component_property='value')
)
def update_city_selected(input_value):
    return f'You selected: {input_value}'