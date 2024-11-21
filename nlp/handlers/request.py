from pydantic import BaseModel

class Req(BaseModel):
    symptoms: list[str]
    input: str
    isTime: bool = False
    isMedicine: bool = False
