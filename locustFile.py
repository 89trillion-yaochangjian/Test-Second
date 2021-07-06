from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    @task
    def Calculator(self):
        self.client.get("CalculationCtrl?str=1+2*3+(4*5+6)*7")