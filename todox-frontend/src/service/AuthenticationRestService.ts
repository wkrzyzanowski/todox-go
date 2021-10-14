import AxiosRepository from '@/plugins/axios/axios';
import { AxiosPromise } from 'axios';

export default class AuthenticationRestService {
  static login(user: string, pass: string): AxiosPromise {
    return AxiosRepository.post('/authentication/login', {
      username: user,
      password: pass,
    });
  }
}
