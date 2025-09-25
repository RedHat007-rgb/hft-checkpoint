import { Module } from '@nestjs/common';

import { TickerService } from './ticker.service';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'TICKER_PACKAGE',
        transport: Transport.GRPC,
        options: {
          package: 'ticker',
          url: 'localhost:50051',
          protoPath: join(__dirname, '../../../packages/proto/ticker.proto'),
        },
      },
    ]),
  ],

  providers: [TickerService],
})
export class TickerModule {}
