import dash
import dash_bootstrap_components as dbc
from dash import html

_LOGO = "http://www.w3.org/2000/svg"


def header():
    def _items():
        _links = []
        for page in dash.page_registry.values():
            if 'order' in page and page['order'] is not None:
                _links.append(
                    dbc.NavItem(dbc.NavLink(page['name'], href=page["relative_path"]))
                )
        return _links

    return dbc.Navbar(
        dbc.Container(
            [
                html.A(
                    dbc.Row(
                        [
                            dbc.Col(html.Img(src=_LOGO, height="30px")),
                            dbc.Col(dbc.NavbarBrand("Upowl", className="ms-2")),
                        ],
                        align="center",
                        className="g-0",
                    ),
                    style={"textDecoration": "none"},
                ),
                dbc.Row(
                    [
                        dbc.NavbarToggler(id="navbar-toggler"),
                        dbc.Collapse(
                            dbc.Nav(
                                _items(),
                                # make sure nav takes up the full width for auto
                                # margin to get applied
                                className="w-100",
                            ),
                            id="navbar-collapse",
                            is_open=False,
                            navbar=True,
                        ),
                    ],
                    # the row should expand to fill the available horizontal space
                    className="flex-grow-1",
                ),
            ]
        ),
        dark=True,
        color="dark",
    )
