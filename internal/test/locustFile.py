from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    @task
    def Calculator(self):
        data = {
            "str":"1*2+3*4-1"
        }
        self.client.post("/CalculationCtrl",data=data)