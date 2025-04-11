FROM python:alpine

WORKDIR /usr/distribution-center-api

COPY ./distribution-center-api/requirements.txt .

RUN pip install --no-cache-dir gunicorn
RUN pip install --no-cache-dir -r requirements.txt

COPY ./distribution-center-api/app ./app

EXPOSE 8001

CMD ["gunicorn", "app:create_app()", "--bind", "0.0.0.0:8001"]