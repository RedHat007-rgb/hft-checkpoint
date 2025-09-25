import {
  TickerAck,
  TickerRequest,
  TickerService as TickerServiceClient,
} from '@app/common';
import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import type { ClientGrpc } from '@nestjs/microservices';
import type { RedisClientType } from 'redis';

@Injectable()
export class TickerService implements OnModuleInit {
  private tickerClient: TickerServiceClient;

  constructor(
    @Inject('TICKER_PACKAGE') private client: ClientGrpc,
    @Inject('REDIS_CLIENT') private redisClient: RedisClientType,
  ) {}

  onModuleInit() {
    this.tickerClient =
      this.client.getService<TickerServiceClient>('TickerService');
  }

  async subscribe(tickerRequest: TickerRequest): Promise<TickerAck> {
    console.log(tickerRequest.symbol);

    const channel = 'binance.' + tickerRequest.symbol + '.ticker';
    await this.redisClient.subscribe(channel, (data) => {
      console.log(data);
    });
    return this.tickerClient.Subscribe(tickerRequest);
  }
}
