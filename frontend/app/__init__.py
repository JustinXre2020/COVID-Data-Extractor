from flask import Flask, render_template, request
import requests
import json


app = Flask(__name__)


@app.route("/")
def homepage():
    return render_template("index.html")


@app.route('/submit', methods=['POST'])
def submitForm():
    operation = request.form['operation']
    table = request.form['table']
    ColumnOrValues = request.form['ColumnOrValues']
    conditions = request.form['conditions']

    if not operation:
        return render_template("index.html", message='Please enter required fields!')

    param = {
        'dataSource': table,
        'updates': ColumnOrValues,
        'condition': conditions
    }

    conValues = {col.split(",")[0].strip(): col.split(",")[1].strip() for col in ColumnOrValues.split(";")} if ColumnOrValues else None
    conditionValues = {condition.split(",")[0].strip(): condition.split(",")[1].strip() for condition in conditions.split(";")} if conditions else None
    print(f"conditions are {conditionValues}")
    param['updates'] = json.dumps(conValues)
    param['condition'] = json.dumps(conditionValues)


    print(f"Operation: {operation}")
    if operation == 'advancequery_1' or operation == 'advancequery_2':
        r = requests.get(f'http://10.195.215.63:3000/v1/api/{operation}')
        data = r.json()
        print(f"Result: {data}")
        if data['code'] == 0:
            return render_template("index.html", message=data['message'])

        sent = json.loads(data['data']['ResponseResult'])
        return render_template("success.html", jsonData=sent)
    elif operation == 'read':
        r = requests.post(f'http://10.195.215.63:3000/v1/api/{operation}',data=param)
        data = r.json()
        print(f"Result: {data}")
        if data['code'] == 0:
            return render_template("index.html", message=data['message'])
            
        sent = json.loads(data['data']['ResponseResult'])
        pan=json.dumps(data)[9:10]
        if pan=='1':
            return render_template("success.html", jsonData=sent)
        else :
            return render_template("success.html")     
    else:
        r = requests.post(f'http://10.195.215.63:3000/v1/api/{operation}', data=param)
        res = r.json()
        print(f"Result: {res}")
        if res['code'] == 0:
            return render_template("index.html", message=res['message'])
        return render_template("success.html")
