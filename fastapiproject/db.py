from motor.motor_asyncio import AsyncIOMotorClient, AsyncIOMotorDatabase


async def connect_to_movies_db() -> AsyncIOMotorDatabase:
    client = AsyncIOMotorClient('mongodb://azfastapiserver:jaWquzgp9MRClHHgCG58wAOs5vlz8azjsu7tpbvc48l6Gj3D4YKdnQGdcSD6ECkCOCkjFkkkdGgWACDbTUTJHg%3D%3D@azfastapiserver.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@azfastapiserver@')
    try:
        db = client['movie_db']
        yield db
    finally:
        client.close()