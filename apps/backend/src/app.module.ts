import { Module } from '@nestjs/common';
import { TickerModule } from './ticker/ticker.module';
import { RedisModule } from './redis/redis.module';

@Module({
  imports: [TickerModule, RedisModule],
  controllers: [],
  providers: [],
})
export class AppModule {}
