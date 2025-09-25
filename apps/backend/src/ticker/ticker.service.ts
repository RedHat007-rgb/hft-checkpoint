import {
  TickerAck,
  TickerRequest,
  TickerService as TickerServiceClient, // <-- rename here
} from '@app/common';
import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import type { ClientGrpc } from '@nestjs/microservices';

@Injectable()
export class TickerService implements OnModuleInit {
  private tickerClient: TickerServiceClient;

  constructor(@Inject('TICKER_PACKAGE') private client: ClientGrpc) {}

  onModuleInit() {
    this.tickerClient =
      this.client.getService<TickerServiceClient>('TickerService');
  }

  subscribe(symbol: TickerRequest): Promise<TickerAck> {
    return this.tickerClient.Subscribe(symbol);
  }
}
