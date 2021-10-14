import AxiosRepository from '@/plugins/axios/axios';
import { AxiosPromise } from 'axios';

export default class HelloRepository {
  static getHello(): AxiosPromise {
    return AxiosRepository.get('/hello');
  }
}
