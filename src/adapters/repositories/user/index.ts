import { User } from '../../../domain/entities/user/user'
import { Email } from '../../../domain/models/email'

export type PaginationOptions = {
  limit?: number
  cursor?: string
}

export type UserRepository = {
  getById(id: string): Promise<User>
  getByEmail(email: Email): Promise<User>
  list(): Promise<{ collection: User[]; cursor?: string }>
  persist(user: User): Promise<void>
}
