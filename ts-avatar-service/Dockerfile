FROM python:3.10-bullseye

RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

RUN apt-get update && apt-get install -y cmake libgl1

RUN mkdir -p /app
ADD . /app/
WORKDIR /app

RUN pip install -r requirements.txt

CMD ["gunicorn", "-w", "4", "-b", "0.0.0.0:17001", "app:app"]

EXPOSE 17001
