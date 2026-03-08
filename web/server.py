import subprocess
from flask import Flask, request, render_template, send_file, jsonify

app = Flask(__name__)
INPUT_FILE = "../input.txt"
OUTPUT_FILE = "./output.txt"
TEMP_FILE = '../temp.txt'

@app.route('/')
def index():
    with open(INPUT_FILE, 'r') as f:
        input_content = f.read()
    return render_template('index.html', input_content=input_content)

@app.route('/run', methods=['POST'])
def run_binary():
    binary_path = request.form.get('binary_path')
    board = request.form.get('board')
    
    with open(TEMP_FILE, 'w', newline='\n') as t:
        board = board.replace('\r\n', '\n')
        t.write(board)
        t.close()

        try:
            subprocess.run([binary_path, "--file", TEMP_FILE, "--amount", "1"], check=True)

            return send_file(OUTPUT_FILE, as_attachment=False)
        except subprocess.CalledProcessError as e:
            return jsonify({"error": f"Error while executing binary: {str(e)}"}), 500

if __name__ == '__main__':
    app.run(debug=True)
