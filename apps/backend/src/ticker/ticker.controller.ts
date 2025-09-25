import { Body, Controller, Get } from '@nestjs/common';
import { TickerService } from './ticker.service';
import { TickerAck, TickerRequest } from '@app/common';

@Controller('ticker')
export class TickerController {
  constructor(private tickerService: TickerService) {}

  @Get()
  subscribe(@Body() symbol: TickerRequest): Promise<TickerAck> {
    console.log(symbol);
    console.log('in controller');
    return this.tickerService.subscribe(symbol);
  }
}
