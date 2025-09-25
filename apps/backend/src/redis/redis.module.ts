import { Global, Module } from '@nestjs/common';
import { createClient, RedisClientType } from 'redis';
@Global()
@Module({
  providers: [
    {
      provide: 'REDIS_CLIENT',
      useFactory: async () => {
        const client: RedisClientType = createClient({
          url: 'redis://localhost:6379',
        });
        client.on('error', (err) => console.error('Redis client error', err));
        await client.connect();
        return client;
      },
    },
  ],
  exports: ['REDIS_CLIENT'],
})
export class RedisModule {}
