from flask import Flask, render_template

app = Flask(__name__)

@app.route('/')
def index():
    # Simulatie user State
    user_logged_in = False

    # Beschikbare acties voor user:
    actions = [
        {"name": "Login", "endpoint": "/login"},
        {"name": "Register", "endpoint": "/register"}
    ]

    if user_logged_in:
        actions.extend([
            {"name": "Add Task", "endpoint": "/add_task"},
            {"name": "View Tasks", "endpoint": "/view_tasks"},
            {"name": "Logout", "endpoint": "/logout"},
            {"name": "Edit Task", "endpoint": "/edit_task"},
            {"name": "Delete Task", "endpoint": "/delete_task"}
        ])

    return render_template('index.html', actions=actions)

if __name__ == '__main__':
    app.run(debug=True)
