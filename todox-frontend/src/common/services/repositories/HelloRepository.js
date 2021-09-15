import AbstractRepository from '@/common/services/repositories/_AbstractRepository';

export default class HelloRepository {
  static getHello() {
    return AbstractRepository.get('/hello');
  }
}
